package model

import "gorm.io/gorm"

// Device represents a mobile device for app testing.
type Device struct {
	gorm.Model
	Type       string `json:"type" gorm:"size:255"` // Phone, Tablet
	Mod        string `json:"model" gorm:"size:255"`
	OS         string `json:"os" gorm:"size:255"`         // Operating System
	Resolution string `json:"resolution" gorm:"size:100"` // Screen Resolution
}
