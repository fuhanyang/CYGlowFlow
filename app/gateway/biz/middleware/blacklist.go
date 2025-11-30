package middleware

import (
	"context"
	"strconv"

	"time"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/fuhanyang/CYGlowFlow/app/gateway/biz/dal/mysql/query"
)

// BlacklistMiddleware 黑名单拦截中间件
func BlacklistMiddleware() app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		clientIP := c.ClientIP()
		var id string
		// 2. 检查 UserID 黑名单 (逻辑同上)
		if v, exists := c.Get("current_user_id"); exists {
			if userID, ok := v.(int64); ok {
				_ = userID
				id = strconv.Itoa(int(userID))
				// ... 同样逻辑
			}
		}
		if id == "" {
			c.JSON(consts.StatusForbidden, utils.H{"code": 403, "msg": "Access Denied:user id is empty"})
			c.Abort()
			return
		}
		// 1. 检查 IP 黑名单
		bl := query.Q.GatewayBlacklist

		// 优化为单次查询：(Type=1 AND Value=IP AND Enable=1) AND (ExpireAt IS NULL OR ExpireAt > NOW)
		count, _ := bl.WithContext(ctx).
			Where(
				bl.WithContext(ctx).Where(
					bl.Type.Eq(1),
					bl.Value.Eq(clientIP),
					bl.Enable.Eq(1),
				).Or(
					bl.Type.Eq(2),
					bl.Value.Eq(id),
					bl.Enable.Eq(1),
				),
				bl.ExpireAt.IsNull()).Or(bl.ExpireAt.Gt(time.Now())).
			Count()

		if count > 0 {
			c.JSON(consts.StatusForbidden, utils.H{"code": 403, "msg": "Access Denied: IP Blacklisted"})
			c.Abort()
			return
		}
		c.Next(ctx)
	}
}
