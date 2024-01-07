package model

import "gorm.io/gorm"

// Template represents a template for test cases or reports.
type Template struct {
	gorm.Model
	Name         string `json:"name" gorm:"size:255"`
	Content      string `json:"content" gorm:"type:text"`
	TemplateType string `json:"template_type" gorm:"size:100"` // 'testcase', 'report', etc.
}
