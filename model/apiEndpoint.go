package model

import "gorm.io/gorm"

type APIEndpoint struct {
	gorm.Model
	URL    string `json:"url" gorm:"size:500"`
	Method string `json:"method" gorm:"size:10"`
}
