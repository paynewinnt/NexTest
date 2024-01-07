package model

import (
	"gorm.io/gorm"
	"time"
)

type Log struct {
	gorm.Model
	Timestamp time.Time `json:"timestamp"`
	Message   string    `json:"message" gorm:"size:1000"`
}
