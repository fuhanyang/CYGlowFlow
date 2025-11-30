package middleware

import (
	"context"
	"strings"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	commonUtils "github.com/fuhanyang/CYGlowFlow/common/utils"
)

// AuthMiddleware 鉴权中间件
func AuthMiddleware() app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		authHeader := string(c.GetHeader("Authorization"))
		if authHeader == "" {
			c.JSON(consts.StatusUnauthorized, utils.H{"code": 401, "msg": "Missing Authorization Header"})
			c.Abort()
			return
		}

		parts := strings.SplitN(authHeader, " ", 2)
		if len(parts) != 2 || parts[0] != "Bearer" {
			c.JSON(consts.StatusUnauthorized, utils.H{"code": 401, "msg": "Invalid Authorization Header Format"})
			c.Abort()
			return
		}

		tokenString := parts[1]
		claims, err := commonUtils.ParseToken(tokenString)
		if err != nil {
			c.JSON(consts.StatusUnauthorized, utils.H{"code": 401, "msg": "Invalid or Expired Token"})
			c.Abort()
			return
		}

		// 将 UserID 存入上下文，供后续 Handler 使用
		c.Set("current_user_id", claims.UserID)

		c.Next(ctx)
	}
}
