package middleware

import (
	"context"
	"sync"
	"time"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

// LimitConfig 限流配置
type LimitConfig struct {
	Capacity int // 令牌桶容量
	Rate     int // 每秒产生令牌数
}

// TokenBucket 简单的令牌桶实现
type TokenBucket struct {
	capacity   int       // 容量
	rate       int       // 速率（每秒）
	tokens     float64   // 当前令牌数
	lastRefill time.Time // 上次填充时间
	mu         sync.Mutex
}

// NewTokenBucket 创建一个新的令牌桶
func NewTokenBucket(capacity, rate int) *TokenBucket {
	return &TokenBucket{
		capacity:   capacity,
		rate:       rate,
		tokens:     float64(capacity), // 初始满桶
		lastRefill: time.Now(),
	}
}

// Allow 尝试消耗一个令牌
func (tb *TokenBucket) Allow() bool {
	tb.mu.Lock()
	defer tb.mu.Unlock()

	now := time.Now()
	// 计算时间差（秒）
	elapsed := now.Sub(tb.lastRefill).Seconds()

	// 补充令牌
	tb.tokens = tb.tokens + elapsed*float64(tb.rate)
	if tb.tokens > float64(tb.capacity) {
		tb.tokens = float64(tb.capacity)
	}

	tb.lastRefill = now

	if tb.tokens >= 1.0 {
		tb.tokens -= 1.0
		return true
	}

	return false
}

// RateLimitMiddleware 简单的内存限流中间件
// 生产环境建议使用 Redis + Lua 脚本实现分布式限流
func RateLimitMiddleware(capacity int, qps int) app.HandlerFunc {
	// 使用 sync.Map 存储每个 IP 的限流器
	// key: IP, value: *TokenBucket
	var limiters sync.Map

	return func(ctx context.Context, c *app.RequestContext) {
		clientIP := c.ClientIP()

		// 获取或创建限流器
		v, ok := limiters.Load(clientIP)
		if !ok {
			// 创建一个新的限流器
			newLimiter := NewTokenBucket(capacity, qps)
			v, _ = limiters.LoadOrStore(clientIP, newLimiter)
		}

		limiter := v.(*TokenBucket)

		// 尝试获取 1 个令牌，如果获取不到（桶空了），则拒绝请求
		if !limiter.Allow() {
			hlog.CtxWarnf(ctx, "Rate limit exceeded for IP: %s", clientIP)
			c.JSON(consts.StatusTooManyRequests, utils.H{
				"code": 429,
				"msg":  "Too Many Requests",
			})
			c.Abort()
			return
		}

		c.Next(ctx)
	}
}
