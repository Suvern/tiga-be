package handler

import (
	"github.com/gin-gonic/gin"
	"tiga/model"
	"tiga/model/form"
	"tiga/service"
)

func AdminLogin(c *gin.Context) {
	// username: admin
	// password: PpVlXgzSO8oh06OABDDYVsGk
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
	c.JSON(200, "updatePassword")
}

func UpdatePreference(c *gin.Context) {
	c.JSON(200, "updatePreference")
}
