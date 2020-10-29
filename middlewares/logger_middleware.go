package middlewares

import (
	"fmt"
	"github.com/zzlpeter/dawn-go/libs/log"
	"github.com/zzlpeter/dawn-go/libs/utils"
	"github.com/gin-gonic/gin"
)

// LoggerMiddleware api接口访问纪录日志中间件
func ApiAccessLoggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 本地开发环境不需要登录
		env := utils.Environ{}.Get("APP_ENV")
		if env != "test" && env != "product" {
			c.Next()
			return
		}

		username, _ := c.Cookie("username")
		query := c.Request.URL.Query()
		c.Request.ParseForm()
		form := c.Request.PostForm
		body := c.Request.Body
		method := c.Request.Method
		uri := c.Request.URL.Path
		msg := fmt.Sprintf("用户%v访问%v", username, uri)
		log.LogRecord(c, log.INFO, msg, "method", method, "query", query, "form", form, "body", body)

		c.Next()
	}
}