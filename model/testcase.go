package model

import (
	"gorm.io/gorm"
	"time"
)

// TestCase 表示一个测试用例实体
type TestCase struct {
	gorm.Model            // 内嵌gorm.Model，包含ID, CreatedAt, UpdatedAt等字段
	Title       string    `gorm:"size:255;not null"` // 测试用例标题
	Description string    `gorm:"size:500"`          // 测试用例描述
	Steps       string    `gorm:"type:text"`         // 测试步骤
	Expected    string    `gorm:"type:text"`         // 预期结果
	Actual      string    `gorm:"type:text"`         // 实际结果
	Status      string    `gorm:"size:100"`          // 测试状态（例如 '未执行', '成功', '失败'）
	Priority    string    `gorm:"size:100"`          // 优先级
	ProjectID   uint      // 关联的项目ID
	CreatedBy   uint      // 创建者ID
	AssignedTo  uint      // 指派给的用户ID
	RunDate     time.Time // 测试执行日期
}
