package model

import (
	"gorm.io/gorm"
	"time"
)

// Error represents an error or exception occurred during testing.
type Error struct {
	gorm.Model
	Timestamp  time.Time `json:"timestamp"`
	TestID     uint      `json:"test_id"`
	Message    string    `json:"message" gorm:"size:500"`
	Severity   string    `json:"severity" gorm:"size:100"`     // Error severity level
	StackTrace string    `json:"stack_trace" gorm:"type:text"` // Error stack trace
}
