package model

import "gorm.io/gorm"

type FeatureFlag struct {
	gorm.Model
	Name      string `json:"name" gorm:"size:255"`
	IsEnabled bool   `json:"is_enabled"`
}
