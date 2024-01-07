package model

import "gorm.io/gorm"

// Scenario represents a testing scenario, such as in end-to-end testing.
type Scenario struct {
	gorm.Model
	Name        string `json:"name" gorm:"size:255"`
	Description string `json:"description" gorm:"size:500"`
	Steps       []Step `json:"steps" gorm:"foreignKey:ScenarioID"` // Test Steps
}
