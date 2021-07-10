package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Db *gorm.DB

func init() {
	dsn := ""
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		print("数据库连接失败")
		panic(err)
	}
	Db = db
}
