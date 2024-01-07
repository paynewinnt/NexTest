package model

import (
	"gorm.io/gorm"
	"time"
)

// User 表示一个用户实体，用于用户管理
type User struct {
	gorm.Model           // 内嵌gorm.Model，包含ID, CreatedAt, UpdatedAt等字段
	Username   string    `gorm:"size:255;not null;unique"` // 用户名，唯一标识
	Password   string    `gorm:"size:255;not null"`        // 用户密码
	Email      string    `gorm:"size:255"`                 // 用户邮箱
	IsActive   bool      `gorm:"default:true"`             // 用户是否激活
	LastLogin  time.Time // 最后登录时间
	// 可以根据需求添加更多字段，如用户角色、权限等
}
