package model

import "gorm.io/gorm"

type Preferences struct {
	gorm.Model
	UserID   uint                   `json:"user_id"`
	Settings map[string]interface{} `json:"settings" gorm:"type:json"`
}
