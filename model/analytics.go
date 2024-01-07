package model

import "gorm.io/gorm"

type Analytics struct {
	gorm.Model
	TestID             uint               `json:"test_id"`
	CoverageRate       float64            `json:"coverage_rate"`
	PerformanceMetrics map[string]float64 `json:"performance_metrics" gorm:"type:json"`
}
