package model

import (
	"gorm.io/gorm"
	"time"
)

// AuditLog 表示一个审计日志实体
type AuditLog struct {
	gorm.Model           // 内嵌gorm.Model
	UserID     uint      `gorm:"index"`     // 执行操作的用户ID
	Action     string    `gorm:"size:255"`  // 执行的操作（例如 'login', 'create_test', 'update_project'）
	Details    string    `gorm:"type:text"` // 操作的详细描述
	Timestamp  time.Time // 操作发生的时间戳
	IPAddress  string    `gorm:"size:255"` // 用户的IP地址
	Success    bool      // 操作是否成功
	Entity     string    `gorm:"size:255"` // 涉及的实体（例如 'Project', 'TestCase', 'User'）
	EntityID   uint      // 涉及实体的ID
	// 可以根据需求添加更多字段
}
