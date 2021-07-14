package service

import (
	"tiga/db"
	"tiga/model"
	"tiga/model/dao"
	"tiga/model/form"
	"tiga/util"
)

func AdminLogin(form form.LoginForm) string {
	setting := dao.Setting{}
	db.Db.Limit(1).Find(&setting)
	if setting.Username == form.Username && util.CheckPasswordHash(form.Password, setting.Password) {
		return util.GenerateToken(form.Username)
	} else {
		panic(model.WrongPasswordError)
	}
}

func UpdatePassword(form form.UpdatePasswordForm) {
	// TODO: 需要事务
	setting := dao.Setting{}
	db.Db.Limit(1).Find(&setting)
	if util.CheckPasswordHash(form.OldPassword, setting.Password) {
		setting.Password = util.HashPassword(form.NewPassword)
		db.Db.Save(&setting)
	} else {
		panic(model.WrongPasswordError)
	}
}

func UpdatePreference(preference string) {
	// TODO: 验证JSON格式
	setting := dao.Setting{}
	db.Db.Limit(1).Find(&setting)
	setting.Preference = preference
	db.Db.Save(&setting)
}
