package middlewares

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// Response200 返回成功信息
func Response200(msg string, data interface{}, c *gin.Context)  {
	c.JSON(http.StatusOK, map[string]interface{}{
		"code": http.StatusOK,
		"message": msg,
		"data": data,
	})
	return
}

// ResponseError 其他错误类型返回信息
func ResponseError(code int, msg string, c *gin.Context)  {
	c.JSON(code, map[string]interface{}{
		"code": code,
		"message": msg,
		"data": nil,
	})
	return
}