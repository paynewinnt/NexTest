package middlewares

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"runtime/debug"
)

// RecoveryMiddleware 创建一个错误恢复中间件
func RecoveryMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Defer一个函数来捕获可能的panic
		defer func() {
			if err := recover(); err != nil {
				// 记录堆栈信息
				log.Printf("panic: %v\n%s", err, string(debug.Stack()))

				// 返回错误响应
				c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
					"error": "内部服务器错误",
				})
			}
		}()

		// 继续后续的处理
		c.Next()
	}
}
