package middlewares

import (
	"github.com/gin-gonic/gin"
)

// AuthMiddleware 登录认证中间件
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("SESSION_ID")
		if token == "" {
			ResponseError(401, "请登录", c)
			c.Abort()
			return
		}
		c.Next()
	}
}