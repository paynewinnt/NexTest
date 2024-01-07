package middlewares

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

// 日志等级
const (
	LogLevelDebug   = "DEBUG"
	LogLevelInfo    = "INFO"
	LogLevelWarning = "WARNING"
	LogLevelError   = "ERROR"
)

// 日志等级变量，根据需要设置
var currentLogLevel = LogLevelInfo

// LoggingMiddleware 创建一个日志记录中间件
func LoggingMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 请求开始时间
		startTime := time.Now()

		// 处理请求
		c.Next()

		// 请求结束时间
		endTime := time.Now()

		// 计算处理时间
		latency := endTime.Sub(startTime)

		// 获取请求方式和路径
		method := c.Request.Method
		path := c.Request.URL.Path

		// 获取响应状态
		status := c.Writer.Status()

		// 根据日志等级和状态码决定日志内容
		logLevel := determineLogLevel(status)
		if shouldLogForLevel(logLevel) {
			// 记录日志
			log.Printf("[%s] %s %s %d %v", logLevel, method, path, status, latency)
		}
	}
}

// determineLogLevel 根据响应状态码确定日志等级
func determineLogLevel(status int) string {
	switch {
	case status >= http.StatusInternalServerError:
		return LogLevelError
	case status >= http.StatusBadRequest:
		return LogLevelWarning
	case status == http.StatusNotFound:
		return LogLevelInfo
	default:
		return LogLevelDebug
	}
}

// shouldLogForLevel 确定当前设置的日志等级是否应该记录指定等级的日志
func shouldLogForLevel(level string) bool {
	var levels = map[string]int{
		LogLevelDebug:   1,
		LogLevelInfo:    2,
		LogLevelWarning: 3,
		LogLevelError:   4,
	}

	return levels[level] >= levels[currentLogLevel]
}
