package model

import (
	"gorm.io/gorm"
)

// TestEnvironment 表示一个测试环境实体
type TestEnvironment struct {
	gorm.Model         // 内嵌gorm.Model
	Name        string `gorm:"size:255;not null;unique"` // 环境名称
	Description string `gorm:"size:500"`                 // 环境描述
	Config      string `gorm:"type:text"`                // 环境配置详情
	ProjectID   uint   // 关联的项目ID
	IsActive    bool   `gorm:"default:true"` // 环境是否激活
	Variables   string `gorm:"type:text"`    // 环境变量
	AccessURL   string `gorm:"size:255"`     // 环境访问URL
	// 更多自定义字段
}
