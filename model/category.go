package model

type Category struct {
	ID       uint      `gorm:"primaryKey"`
	Name     string    `gorm:"unique"`
	Articles []Article `gorm:"polymorphic:Category;"`
}
