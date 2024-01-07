package model

import "gorm.io/gorm"

type Dashboard struct {
	gorm.Model
	UserID  uint     `json:"user_id"`
	Layout  string   `json:"layout" gorm:"size:1000"`
	Widgets []string `json:"widgets" gorm:"type:json"`
}
