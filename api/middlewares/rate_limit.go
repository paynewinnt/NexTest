package middlewares

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"
	"net/http"
	"sync"
)

/*
在这个中间件中：
为每个客户端IP地址维护一个 rate.Limiter 实例。
使用客户端IP地址作为限制的依据（您也可以根据其他标识符进行限制）。
如果客户端超过了设定的请求频率（limit），则返回HTTP 429 Too Many Requests状态码。
*/

// 限制器映射，用于存储每个客户端的限制器
var limiterMap = make(map[string]*rate.Limiter)
var mu sync.Mutex

// RateLimitMiddleware 创建一个速率限制中间件
// limit 是每秒允许的请求数
// burst 是限制器的最大突发大小
func RateLimitMiddleware(limit rate.Limit, burst int) gin.HandlerFunc {
	return func(c *gin.Context) {
		clientIP := c.ClientIP()

		mu.Lock()
		limiter, exists := limiterMap[clientIP]
		if !exists {
			limiter = rate.NewLimiter(limit, burst)
			limiterMap[clientIP] = limiter
		}
		mu.Unlock()

		if !limiter.Allow() {
			c.AbortWithStatusJSON(http.StatusTooManyRequests, gin.H{"error": "请求过于频繁"})
			return
		}

		c.Next()
	}
}
