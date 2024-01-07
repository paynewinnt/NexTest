package model

import "gorm.io/gorm"

// Browser represents a web browser for web automation testing.
type Browser struct {
	gorm.Model
	Name    string `json:"name" gorm:"size:255"`
	Version string `json:"version" gorm:"size:100"`
	Engine  string `json:"engine" gorm:"size:100"` // Browser engine like WebKit, Blink
}
