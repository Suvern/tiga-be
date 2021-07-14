package service

import (
	"tiga/db"
	"tiga/model/dao"
	"tiga/util"
)

func AdminLogin(username, password string) {
	var setting = dao.Setting{}
	db.Db.Limit(1).Find(&setting)
	if setting.Username == username && util.CheckPasswordHash(password, setting.Password) {
		// 登录成功
	} else {
		// 用户不存在或密码错误
	}
}

func UpdatePassword(password string) {
	// TODO: 需要事务
	var setting = dao.Setting{}
	db.Db.Limit(1).Find(&setting)
	setting.Password = util.HashPassword(password)
	db.Db.Save(&setting)
}

func UpdatePreference(preference string) {
	var setting = dao.Setting{}
	db.Db.Limit(1).Find(&setting)
	// TODO: 验证JSON格式
	setting.Preference = preference
	db.Db.Save(&setting)
}
