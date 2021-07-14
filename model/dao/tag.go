package dao

type Tag struct {
	ID       uint      `gorm:"primaryKey"`
	Name     string    `gorm:"unique"`
	Articles []Article `gorm:"many2many:article_tag"`
}
