package middlewares

import (
	"github.com/gin-gonic/gin"
)

// SecurityHeadersMiddleware 创建一个设置安全头的中间件
func SecurityHeadersMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 设置安全相关的头部
		c.Header("X-Frame-Options", "DENY")                      // 防止点击劫持
		c.Header("X-Content-Type-Options", "nosniff")            // 防止MIME类型混淆攻击
		c.Header("X-XSS-Protection", "1; mode=block")            // 启用XSS保护
		c.Header("Content-Security-Policy", "script-src 'self'") // 设置内容安全策略

		// 继续后续的处理
		c.Next()
	}
}
