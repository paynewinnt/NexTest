package model

import (
	"gorm.io/gorm"
	"time"
)

// Asset 表示测试过程中使用的资产，如测试脚本、配置文件等
type Asset struct {
	gorm.Model            // 内嵌gorm.Model，包含ID, CreatedAt, UpdatedAt等字段
	Name        string    `gorm:"size:255;not null;unique"` // 资产名称
	Type        string    `gorm:"size:100;not null"`        // 资产类型（例如 'script', 'config' 等）
	Content     string    `gorm:"type:text"`                // 资产内容或路径
	Description string    `gorm:"size:500"`                 // 资产描述
	OwnerID     uint      `gorm:"not null"`                 // 资产拥有者ID
	ExpiresAt   time.Time // 资产过期时间（可选）
}
