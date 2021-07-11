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
		article.PUT("/:id", handler.UpdateArticle)
		article.DELETE("/:id", handler.DeleteArticle)
	}

	category := r.Group("category")
	{
		category.GET("/", handler.GetCategoryList)
		category.POST("/", handler.CreateCategory)
		category.PUT("/", handler.UpdateCategory)
		category.DELETE("/", handler.DeleteCategory)
	}

	tag := r.Group("tag")
	{
		tag.GET("/", handler.GetTagList)
		tag.POST("/", handler.CreateTag)
		tag.DELETE("/", handler.DeleteTag)
	}
}
