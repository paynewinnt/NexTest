package model

import "gorm.io/gorm"

// Validation represents a validation rule or check.
type Validation struct {
	gorm.Model
	Rule    string `json:"rule" gorm:"size:500"`
	Message string `json:"message" gorm:"size:500"`
	TestID  uint   `json:"test_id"` // Associated Test ID
}
