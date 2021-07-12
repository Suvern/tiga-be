package handler

import "github.com/gin-gonic/gin"

func AdminLogin(c *gin.Context) {
	// username: admin
	// password: PpVlXgzSO8oh06OABDDYVsGk
	c.JSON(200, "adminLogin")
}

func UpdatePassword(c *gin.Context) {
	c.JSON(200, "updatePassword")
}

func UpdatePreference(c *gin.Context) {
	c.JSON(200, "updatePreference")
}
