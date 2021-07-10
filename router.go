package main

import "github.com/gin-gonic/gin"

func InitialRouter(r *gin.Engine) {
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, "hello my son.")
	})
}
