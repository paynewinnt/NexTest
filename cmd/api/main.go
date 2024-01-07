package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
)

func main() {
	// 创建一个Gin路由器实例
	router := gin.Default()

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
}

// 可以根据需要添加其他辅助函数或中间件
