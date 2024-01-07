package model

import (
	"gorm.io/gorm"
	"time"
)

// Session represents a user session in the system.
type Session struct {
	gorm.Model
	UserID    uint      `json:"user_id"`
	StartedAt time.Time `json:"started_at"`
	ExpiresAt time.Time `json:"expires_at"`
	IP        string    `json:"ip" gorm:"size:100"` // IP Address
	DeviceID  uint      `json:"device_id"`          // Associated Device ID
}
