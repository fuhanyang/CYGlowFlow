package middleware

import (
	"context"
	"time"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/fuhanyang/CYGlowFlow/app/gateway/biz/dal/mysql/model"
	"github.com/fuhanyang/CYGlowFlow/app/gateway/biz/dal/mysql/query"
	"github.com/fuhanyang/CYGlowFlow/common/routine"
)

var accessLogBatcher *routine.BatchProcessor[*model.GatewayAccessLog]

func init() {
	// 初始化日志批量处理器
	// 策略：积攒 100 条日志，或者每隔 3 秒，触发一次批量写入
	accessLogBatcher = routine.NewBatchProcessor(100, 3*time.Second, func(logs []*model.GatewayAccessLog) {
		if len(logs) == 0 {
			return
		}
		// 使用 CreateInBatches 进行批量插入
		// 注意：这里使用 Background context，因为原请求早已结束
		err := query.Q.GatewayAccessLog.WithContext(context.Background()).CreateInBatches(logs, len(logs))
		if err != nil {
			hlog.Errorf("Failed to batch write access logs: %v", err)
		}
	})
}

// AccessLogMiddleware 访问日志中间件
func AccessLogMiddleware() app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		start := time.Now()

		// 继续处理请求
		c.Next(ctx)

		end := time.Now()
		latency := end.Sub(start).Milliseconds()

		// 1. 提取关键字段 (在主协程中快速完成，避免使用 c.Copy() 带来的深拷贝开销)
		var userID uint64
		if v, exists := c.Get("current_user_id"); exists {
			if id, ok := v.(int64); ok {
				userID = uint64(id)
			}
		}

		// string(bytes) 会发生内存拷贝，但比 c.Copy() 整个 Context 轻量得多
		traceID := string(c.GetHeader("X-Trace-ID"))
		clientIP := c.ClientIP()
		method := string(c.Method())
		path := string(c.Path())
		queryStr := string(c.QueryArgs().QueryString())
		statusCode := int32(c.Response.StatusCode())

		logEntry := &model.GatewayAccessLog{
			TraceID:    &traceID,
			UserID:     &userID,
			IP:         &clientIP,
			Method:     &method,
			Path:       &path,
			Query:      &queryStr,
			StatusCode: &statusCode,
			Latency:    &latency,
		}

		// 仅在错误时记录 ErrorMsg，节省空间
		if statusCode >= 400 {
			errMsg := c.Errors.String()
			logEntry.ErrorMsg = &errMsg
		}

		// 2. 放入批量处理器 (非阻塞，极快)
		accessLogBatcher.Add(logEntry)
	}
}
