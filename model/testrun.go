package model

import (
	"gorm.io/gorm"
	"time"
)

// TestRun 表示一个测试运行实体
type TestRun struct {
	gorm.Model                 // 内嵌gorm.Model
	Name          string       `gorm:"size:255;not null"` // 测试运行名称
	Description   string       `gorm:"size:500"`          // 测试运行描述
	TestSuiteID   uint         // 关联的测试套件ID
	Status        string       `gorm:"size:100"` // 测试运行状态
	StartDate     time.Time    // 测试开始时间
	EndDate       time.Time    // 测试结束时间
	Results       []TestResult // 关联的测试结果列表
	TriggeredBy   uint         // 触发测试运行的用户ID
	Environment   string       `gorm:"size:255"`  // 测试环境标识
	ExecutionLog  string       `gorm:"type:text"` // 执行日志
	Configuration string       `gorm:"type:text"` // 测试配置
	// 可以根据需求添加更多字段
}
