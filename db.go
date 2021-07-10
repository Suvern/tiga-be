package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Db *gorm.DB

func init() {
	dbConfig := Config.GetStringMapString("db")
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/tiga?charset=utf8mb4&parseTime=True&loc=Local", dbConfig["username"], dbConfig["password"], dbConfig["host"], dbConfig["port"])
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		print("数据库连接失败")
		panic(err)
	}
	Db = db
}
