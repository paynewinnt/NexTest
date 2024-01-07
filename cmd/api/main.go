package main

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/paynewinnt/NexTest/api/middlewares"
	"github.com/paynewinnt/NexTest/config"
	"log"
	"net/http"
	"os"
)

func main() {
	// 创建一个Gin路由器实例
	router := gin.Default()
	router.Use(middlewares.AuthMiddleware())
	router.Use(middlewares.RecoveryMiddleware())
	router.Use(middlewares.RateLimitMiddleware(5, 10))
	router.Use(middlewares.MetricsMiddleware())
	router.Use(middlewares.LoggingMiddleware())
	// 应用中间件，允许JSON和XML内容类型
	router.Use(middlewares.ContentTypeMiddleware("application/json", "application/xml"))
	router.Use(middlewares.SecurityHeadersMiddleware())
	router.Use(middlewares.SessionMiddleware(config.SessionSecret))
	router.Use(middlewares.TraceMiddleware())

	// 设置基本路由
	setupRoutes(router)

	// 从环境变量获取端口，如果未设置，则默认为8080
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// 启动Gin服务器
	if err := router.Run(":" + port); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}

// setupRoutes 配置路由和处理函数
func setupRoutes(router *gin.Engine) {
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "NexTest API is up and running!",
		})
	})

	// 添加更多的路由和处理函数
	// 例如: router.GET("/api/testing", apiTestingHandler)
	// 登录路由
	router.POST("/login", func(c *gin.Context) {
		session := sessions.Default(c)
		// 这里应有登录逻辑，例如检查用户名和密码
		// 假设用户登录成功，设置会话
		session.Set("user", "username")
		session.Save()
		c.JSON(200, gin.H{"message": "登录成功"})
	})

	// 受保护的路由
	router.GET("/protected", func(c *gin.Context) {
		session := sessions.Default(c)
		user := session.Get("user")
		if user == nil {
			c.JSON(403, gin.H{"message": "未授权"})
			return
		}
		c.JSON(200, gin.H{"message": "欢迎 " + user.(string)})
	})

	// 登出路由
	router.GET("/logout", func(c *gin.Context) {
		session := sessions.Default(c)
		session.Clear()
		session.Save()
		c.JSON(200, gin.H{"message": "登出成功"})
	})
}

// 可以根据需要添加其他辅助函数或中间件
