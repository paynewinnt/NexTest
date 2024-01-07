package model

import (
	"gorm.io/gorm"
)

// Config 表示应用程序的配置实体
type Config struct {
	gorm.Model         // 内嵌gorm.Model
	Key         string `gorm:"size:255;not null;unique"` // 配置键
	Value       string `gorm:"type:text"`                // 配置值
	Description string `gorm:"size:500"`                 // 配置描述

	// 可以根据需求添加更多字段，如配置类型、作用域等
	Type       string `gorm:"size:255"`     // 配置类型（例如 'string', 'integer', 'boolean'）
	Scope      string `gorm:"size:255"`     // 配置作用域（例如 'global', 'user', 'project'）
	IsEditable bool   `gorm:"default:true"` // 是否可编辑
}
