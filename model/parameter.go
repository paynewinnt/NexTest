package model

import "gorm.io/gorm"

type Parameter struct {
	gorm.Model
	Name  string `json:"name" gorm:"size:255"`
	Value string `json:"value" gorm:"size:500"`
}
