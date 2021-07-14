package main

import (
	"github.com/gin-gonic/gin"
	"tiga/middleware"
)

func main() {
	r := gin.Default()
	r.Use(middleware.Recover)
	InitialRouter(r)
	_ = r.Run()
}
