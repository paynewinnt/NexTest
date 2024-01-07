package model

import "gorm.io/gorm"

// Tag represents a tag that can be assigned to test cases, projects, etc.
type Tag struct {
	gorm.Model
	Name string `json:"name" gorm:"size:255"`
	Type string `json:"type" gorm:"size:100"` // Tag type (e.g., 'project', 'testcase')
}
