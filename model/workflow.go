package model

import "gorm.io/gorm"

// Workflow represents a workflow in the testing or development process.
type Workflow struct {
	gorm.Model
	Name        string `json:"name" gorm:"size:255"`
	Description string `json:"description" gorm:"size:500"`
	Steps       []Step `json:"steps" gorm:"foreignKey:WorkflowID"` // Workflow Steps
}
