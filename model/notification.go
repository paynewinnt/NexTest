package model

import (
	"gorm.io/gorm"
	"time"
)

// Notification 表示一个通知实体
type Notification struct {
	gorm.Model            // 内嵌gorm.Model
	Title       string    `gorm:"size:255;not null"`  // 通知标题
	Message     string    `gorm:"type:text;not null"` // 通知内容
	Type        string    `gorm:"size:100"`           // 通知类型（例如 'Email', 'SMS', 'System'）
	RecipientID uint      // 接收者ID
	SentDate    time.Time // 发送日期
	Read        bool      `gorm:"default:false"` // 是否已读
	Priority    string    `gorm:"size:100"`      // 通知优先级
	// 更多自定义字段
}
