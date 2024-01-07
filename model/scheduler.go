package model

import (
	"gorm.io/gorm"
	"time"
)

type Scheduler struct {
	gorm.Model
	TestID    uint          `json:"test_id"`
	StartTime time.Time     `json:"start_time"`
	Interval  time.Duration `json:"interval"`
}
