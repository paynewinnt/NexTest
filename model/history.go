package model

import (
	"gorm.io/gorm"
	"time"
)

// History represents a record of actions performed in the system.
type History struct {
	gorm.Model
	Action    string    `json:"action" gorm:"size:255"`
	Timestamp time.Time `json:"timestamp"`
	UserID    uint      `json:"user_id"`
	Details   string    `json:"details" gorm:"type:text"` // Additional action details
}
