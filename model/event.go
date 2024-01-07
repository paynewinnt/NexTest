package model

import (
	"gorm.io/gorm"
	"time"
)

type Event struct {
	gorm.Model
	Timestamp time.Time `json:"timestamp"`
	Type      string    `json:"type" gorm:"size:255"`
	Details   string    `json:"details" gorm:"size:1000"`
}
