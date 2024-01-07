package model

import "gorm.io/gorm"

type Auth struct {
	gorm.Model
	Token     string `json:"token" gorm:"size:255"`
	SessionID string `json:"session_id" gorm:"size:255"`
}
