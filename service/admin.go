package service

import (
	"tiga/db"
	"tiga/model"
	"tiga/model/dao"
	"tiga/model/form"
	"tiga/util"
)

func AdminLogin(form form.LoginForm) string {
	var setting = dao.Setting{}
	db.Db.Limit(1).Find(&setting)
	if setting.Username == form.Username && util.CheckPasswordHash(form.Password, setting.Password) {
		return util.GenerateToken(form.Username)
	} else {
		panic(model.WrongPasswordError)
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
