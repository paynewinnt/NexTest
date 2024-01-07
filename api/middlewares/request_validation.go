package middlewares

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

/*
// 使用实例
router := gin.Default()
// 自定义的请求验证逻辑
func customValidation(c *gin.Context) bool {
    // 例如，验证请求体是否为JSON格式
    if c.ContentType() != "application/json" {
        return false
    }
    // 进一步的验证逻辑...
    return true
}
router.Use(middlewares.RequestValidationMiddleware(customValidation))
// 定义路由
router.POST("/some-api", func(c *gin.Context) {
    // 处理逻辑
})

*/

// RequestValidationMiddleware 创建一个请求验证中间件
func RequestValidationMiddleware(validateFunc func(*gin.Context) bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		if !validateFunc(c) {
			// 如果验证失败，中断请求并返回错误
			c.JSON(http.StatusBadRequest, gin.H{"error": "请求数据格式不正确"})
			c.Abort()
			return
		}

		// 继续后续的处理
		c.Next()
	}
}
