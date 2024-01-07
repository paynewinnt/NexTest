package model

import (
	"gorm.io/gorm"
)

// Integration 表示一个第三方集成实体
type Integration struct {
	gorm.Model         // 内嵌gorm.Model
	Name        string `gorm:"size:255;not null;unique"` // 集成名称
	Description string `gorm:"size:500"`                 // 集成描述
	Type        string `gorm:"size:100"`                 // 集成类型（例如 'CI/CD', 'VersionControl', 'Monitoring'）
	Config      string `gorm:"type:text"`                // 集成配置信息（通常为JSON或者YAML）

	// 可以根据需求添加更多字段
	IsActive  bool `gorm:"default:true"` // 集成是否激活
	ProjectID uint // 如果集成特定于项目，则关联的项目ID
	// 更多自定义字段
}
