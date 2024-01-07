package model

import (
	"gorm.io/gorm"
	"time"
)

// TestResult 表示一个测试结果实体
type TestResult struct {
	gorm.Model              // 内嵌gorm.Model
	Status        string    `gorm:"size:100;not null"` // 测试结果状态
	Details       string    `gorm:"type:text"`         // 测试结果详细描述
	TestCaseID    uint      // 关联的测试用例ID
	ProjectID     uint      // 关联的项目ID
	ExecutedBy    uint      // 执行者ID
	ExecutionDate time.Time // 测试执行日期

	// 新增字段
	ErrorLog        string `gorm:"type:text"` // 错误日志
	PerformanceData string `gorm:"type:text"` // 性能数据
	ScreenshotURL   string `gorm:"size:255"`  // 测试截图URL
	Environment     string `gorm:"size:255"`  // 测试环境标识
	// 可以继续根据需求添加更多字段
}
