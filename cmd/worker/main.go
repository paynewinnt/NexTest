package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	// 设置信号处理，以便优雅地处理中断信号
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	// 启动后台工作
	go startWorker()

	// 等待信号
	sig := <-sigs
	fmt.Println("接收到信号:", sig)
	fmt.Println("正在停止工作进程...")

	// 在这里执行任何清理工作
	cleanup()

	fmt.Println("工作进程已停止")
}

// startWorker 函数模拟后台工作进程的任务
func startWorker() {
	fmt.Println("工作进程启动，正在执行任务...")
	// 示例：无限循环，模拟持续的工作
	for {
		// 执行一些工作，例如处理队列
		processTask()

		// 等待一段时间
		time.Sleep(1 * time.Second)
	}
}

// processTask 模拟一个工作任务
func processTask() {
	fmt.Println("执行工作任务...")
	// 实际工作任务的逻辑
}

// cleanup 执行任何必要的清理工作
func cleanup() {
	fmt.Println("执行清理工作...")
	// 清理逻辑
}
