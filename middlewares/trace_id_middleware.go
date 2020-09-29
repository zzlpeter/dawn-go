package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/zzlpeter/dawn-go/libs/utils"
)

// TraceIDMiddleware 设置trace-ID中间件
func TraceIDMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		traceId := utils.GenUUID()
		c.Writer.Header().Set("X-Request-Id", traceId)
		c.Next()
	}
}