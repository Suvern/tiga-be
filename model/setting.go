package model

import "gorm.io/gorm"

type Setting struct {
	gorm.Model
	Username string `gorm:"unique"`
	Password string
}
