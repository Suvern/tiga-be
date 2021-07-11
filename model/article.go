package model

import "gorm.io/gorm"

type Article struct {
	gorm.Model
	Title        string
	Content      string
	Tags         []Tag `gorm:"many2many:article_tag"`
	CategoryType string
	CategoryID   uint
}
