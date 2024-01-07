package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	fmt.Println("执行复杂测试...")

	// 示例：运行Go的单元测试
	if err := runGoTests(); err != nil {
		fmt.Fprintf(os.Stderr, "测试失败: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("所有测试执行完毕.")
}

// runGoTests 执行Go的单元测试
func runGoTests() error {
	// 使用go test命令执行测试
	cmd := exec.Command("go", "test", "./...")

	// 设置输出，以便我们可以看到测试结果
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	// 运行测试命令
	fmt.Println("运行Go单元测试...")
	if err := cmd.Run(); err != nil {
		return err
	}

	return nil
}

// 您可以扩展此程序以运行不同类型的测试，例如性能测试、端到端测试等。
