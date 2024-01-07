package middlewares

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

// SessionMiddleware 创建一个会话管理中间件
func SessionMiddleware(secret string) gin.HandlerFunc {
	store := cookie.NewStore([]byte(secret))
	return sessions.Sessions("mysession", store)
}
