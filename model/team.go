package model

import "gorm.io/gorm"

type Team struct {
	gorm.Model
	Name    string `json:"name" gorm:"size:255"`
	Members []User `json:"members" gorm:"foreignKey:TeamID"`
}
