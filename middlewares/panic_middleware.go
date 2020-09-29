package middlewares

import "github.com/gin-gonic/gin"

func PanicCatchMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				ResponseError(500, err.(string), c)
				c.Abort()
				return
			}
		}()
		c.Next()
	}
}
