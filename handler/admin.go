package handler

import (
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"tiga/model"
	"tiga/model/form"
	"tiga/service"
)

func AdminLogin(c *gin.Context) {
	loginForm := form.LoginForm{}
	if err := c.ShouldBindJSON(&loginForm); err != nil {
		panic(model.UnexpectedParamsError)
	}
	token := service.AdminLogin(loginForm)
	c.SetCookie(
		"token",
		token,
		3600*24*7,
		"/",
		"localhost",
		false,
		false,
	)
	c.JSON(200, "登录成功")
}

func UpdatePassword(c *gin.Context) {
	updatePasswordForm := form.UpdatePasswordForm{}
	if err := c.ShouldBindJSON(&updatePasswordForm); err != nil {
		panic(model.UnexpectedParamsError)
	}
	service.UpdatePassword(updatePasswordForm)
	c.JSON(201, "操作成功")
}

func UpdatePreference(c *gin.Context) {
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		panic(model.UnexpectedParamsError)
	}
	service.UpdatePreference(string(body))
	c.JSON(201, "操作成功")
}
