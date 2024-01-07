package model

import (
	"gorm.io/gorm"
	"time"
)

// Project 表示一个测试项目实体
type Project struct {
	gorm.Model         // 内嵌gorm.Model
	Name        string `gorm:"size:255;not null;unique"`
	Description string `gorm:"size:500"`
	OwnerID     uint   `gorm:"not null"`
	IsActive    bool   `gorm:"default:true"`
	StartDate   time.Time
	EndDate     time.Time

	// 与 TestCase 的一对多关系
	TestCases []TestCase `gorm:"foreignKey:ProjectID"`

	// 与 TestEnvironment 的一对多关系
	TestEnvironments []TestEnvironment `gorm:"foreignKey:ProjectID"`

	// 与 User 的多对多关系（项目团队成员）
	TeamMembers []*User `gorm:"many2many:project_users;"`
}
