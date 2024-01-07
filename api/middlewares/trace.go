package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// TraceMiddleware 创建一个请求跟踪中间件，为每个请求生成唯一的跟踪ID
func TraceMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 生成一个唯一的请求ID
		traceID := uuid.New().String()

		// 将请求ID设置到Context的Key中，以便在后续处理中使用
		c.Set("TraceID", traceID)

		// 将请求ID添加到响应头中
		c.Header("X-Trace-ID", traceID)

		// 继续后续的处理
		c.Next()
	}
}
