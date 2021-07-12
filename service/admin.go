package service

import (
	"tiga/dao"
	"tiga/model"
	"tiga/util"
)

func AdminLogin(username, password string) {
	var setting = model.Setting{}
	dao.Db.Limit(1).Find(&setting)
	if setting.Username == username && util.CheckPasswordHash(password, setting.Password) {
		// 登录成功
	} else {
		// 用户不存在或密码错误
	}
}

func UpdatePassword(password string) {
	// TODO: 需要事务
	var setting = model.Setting{}
	dao.Db.Limit(1).Find(&setting)
	setting.Password = util.HashPassword(password)
	dao.Db.Save(&setting)
}

func UpdatePreference(preference string) {
	var setting = model.Setting{}
	dao.Db.Limit(1).Find(&setting)
	// TODO: 验证JSON格式
	setting.Preference = preference
	dao.Db.Save(&setting)
}
