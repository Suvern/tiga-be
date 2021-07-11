package main

import (
	"github.com/gin-gonic/gin"
	"tiga/handler"
)

func InitialRouter(r *gin.Engine) {
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, "hello my son.")
	})

	article := r.Group("article")
	{
		article.GET("/", handler.GetArticleList)
		article.GET("/tag/:tag", handler.GetArticleList)
		article.GET("/category/:category", handler.GetArticleList)
		article.GET("/detail/:id", handler.GetArticleDetail)
		article.POST("/", handler.CreateArticle)
		article.PUT("/", handler.UpdateArticle)
		article.DELETE("/", handler.DeleteArticle)
	}
}
