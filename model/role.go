package model

import (
	"gorm.io/gorm"
)

// Role 表示一个用户角色实体
type Role struct {
	gorm.Model               // 内嵌gorm.Model
	Name        string       `gorm:"size:255;not null;unique"`    // 角色名称
	Description string       `gorm:"size:500"`                    // 角色描述
	Permissions []Permission `gorm:"many2many:role_permissions;"` // 关联的权限列表
	DefaultRole bool         `gorm:"default:false"`               // 是否为默认角色
	SystemRole  bool         `gorm:"default:false"`               // 是否为系统预定义角色
	// 可以根据需求添加更多字段，如关联的用户列表等
}
