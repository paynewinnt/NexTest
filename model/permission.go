package model

import (
	"gorm.io/gorm"
)

// Permission 表示一个权限实体
type Permission struct {
	gorm.Model         // 内嵌gorm.Model
	Name        string `gorm:"size:255;not null;unique"` // 权限名称
	Description string `gorm:"size:500"`                 // 权限描述
	Action      string `gorm:"size:255;not null"`        // 关联的动作（例如 'read', 'write', 'execute'）
	Resource    string `gorm:"size:255"`                 // 关联的资源（例如 'project', 'testcase', 'report'）
	Scope       string `gorm:"size:255"`                 // 权限作用范围（例如 'global', 'project', 'personal'）
	// 可以根据需求添加更多字段
}
