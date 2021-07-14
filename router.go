package main

import (
	"github.com/gin-gonic/gin"
	"tiga/handler"
	"tiga/middleware"
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
		article.GET("/:id", handler.GetArticleDetail)
		auth := article.Group("")
		{
			auth.Use(middleware.AuthCookie)
			auth.POST("/", handler.CreateArticle)
			auth.PUT("/:id", handler.UpdateArticle)
			auth.DELETE("/:id", handler.DeleteArticle)
		}
	}

	category := r.Group("category")
	{
		category.GET("/", handler.GetCategoryList)
		auth := category.Group("")
		{
			auth.Use(middleware.AuthCookie)
			auth.POST("/", handler.CreateCategory)
			auth.PUT("/", handler.UpdateCategory)
			auth.DELETE("/", handler.DeleteCategory)
		}
	}

	tag := r.Group("tag")
	{
		tag.GET("/", handler.GetTagList)
		auth := tag.Group("")
		{
			auth.Use(middleware.AuthCookie)
			auth.POST("/", handler.CreateTag)
			auth.DELETE("/", handler.DeleteTag)
		}
	}

	admin := r.Group("admin")
	{
		admin.POST("/login", handler.AdminLogin)
		auth := admin.Group("")
		{
			auth.Use(middleware.AuthCookie)
			auth.PUT("/update/password", handler.UpdatePassword)
			auth.PUT("/update/preference", handler.UpdatePreference)
		}
	}
}
