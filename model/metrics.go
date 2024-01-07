package model

import "gorm.io/gorm"

// Metrics represents performance metrics collected during testing.
type Metrics struct {
	gorm.Model
	TestID       uint    `json:"test_id"`
	MemoryUsage  int     `json:"memory_usage"`  // Memory usage in MB
	CPUUsage     float64 `json:"cpu_usage"`     // CPU usage in percentage
	ResponseTime int     `json:"response_time"` // Response time in milliseconds
	Throughput   int     `json:"throughput"`    // Requests per second
}
