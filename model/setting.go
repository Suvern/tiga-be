package model

import "gorm.io/gorm"

type Setting struct {
	gorm.Model
	Username   string
	Password   string
	Preference string `gorm:"type:json"`
}
