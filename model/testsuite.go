package model

import (
	"gorm.io/gorm"
)

// TestSuite 表示一个测试套件实体
type TestSuite struct {
	gorm.Model         // 内嵌gorm.Model
	Name        string `gorm:"size:255;not null;unique"` // 测试套件名称
	Description string `gorm:"size:500"`                 // 测试套件描述
	ProjectID   uint   // 关联的项目ID
	IsActive    bool   `gorm:"default:true"` // 测试套件是否激活
	TestCaseIDs []uint // 关联的测试用例ID列表
	// 可以根据需求添加更多字段，如执行计划、依赖关系等
}
