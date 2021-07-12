package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"tiga/model"
	"tiga/util"
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
	if err := db.AutoMigrate(&model.Article{}, &model.Tag{}, &model.Category{}, &model.Setting{}); err != nil {
		print("自动建表失败")
		panic(err)
	}
	Db = db
	initSetting()
}

func initSetting() {
	var setting model.Setting
	Db.Limit(1).Find(&setting)
	if setting.Username == "" {
		username := "admin"
		password := util.GenerateRandomPassword(24)
		setting.Username = username
		setting.Password = util.HashPassword(password)
		setting.Preference = "{}"
		result := Db.Save(&setting)
		if result.Error != nil {
			print("初始化密码失败")
			panic(result.Error)
		}
		fmt.Println("----------------------------------------")
		fmt.Printf("初始化账号：%s\n", username)
		fmt.Printf("初始化密码： %s\n", password)
		fmt.Println("----------------------------------------")
	}
}
