package middlewares

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

// ContentTypeMiddleware 检查请求的Content-Type是否在允许的类型列表中
func ContentTypeMiddleware(allowedTypes ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		// 获取请求的Content-Type
		reqContentType := c.GetHeader("Content-Type")

		// 检查Content-Type是否在允许的列表中
		isValidType := false
		for _, t := range allowedTypes {
			if strings.Contains(reqContentType, t) {
				isValidType = true
				break
			}
		}

		// 如果Content-Type不在允许的类型中，返回错误响应
		if !isValidType {
			c.JSON(http.StatusUnsupportedMediaType, gin.H{
				"error": "无效的Content-Type",
			})
			c.Abort()
			return
		}

		c.Next()
	}
}
