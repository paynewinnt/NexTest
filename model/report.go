package model

import "gorm.io/gorm"

type Report struct {
	gorm.Model
	TestID  uint   `json:"test_id"`
	Content string `json:"content" gorm:"type:text"`
}
