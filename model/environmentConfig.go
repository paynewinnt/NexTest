package model

import "gorm.io/gorm"

type EnvironmentConfig struct {
	gorm.Model
	Name          string            `json:"name" gorm:"size:255"`
	Variables     map[string]string `json:"variables" gorm:"type:json"`
	ServerAddress string            `json:"server_address" gorm:"size:500"`
}
