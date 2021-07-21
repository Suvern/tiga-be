package handler

import (
	"github.com/gin-gonic/gin"
	"tiga/model"
	"tiga/model/form"
	"tiga/service"
)

func GetArticleListAll(c *gin.Context) {
	articleQueryForm := form.ArticleQueryForm{}
	if err := c.ShouldBindQuery(&articleQueryForm); err != nil {
		panic(model.UnexpectedParamsError)
	}
	articles := service.GetArticleList(articleQueryForm)
	c.JSON(200, articles)
}

func GetArticleDetail(c *gin.Context) {
	articleUriForm := form.ArticleIDForm{}
	if err := c.ShouldBindUri(&articleUriForm); err != nil {
		panic(model.UnexpectedParamsError)
	}
	article := service.GetArticleDetail(articleUriForm.ID)
	c.JSON(200, article)
}

func CreateArticle(c *gin.Context) {
	articleJsonForm := form.ArticleJsonForm{}
	if err := c.ShouldBindJSON(&articleJsonForm); err != nil {
		panic(model.UnexpectedParamsError)
	}
	service.CreateArticle(articleJsonForm)
	c.JSON(201, "操作成功")
}

func UpdateArticle(c *gin.Context) {
	articleJsonForm := form.ArticleJsonForm{}
	if err := c.ShouldBindJSON(&articleJsonForm); err != nil {
		panic(model.UnexpectedParamsError)
	}

	articleIDForm := form.ArticleIDForm{}
	if err := c.ShouldBindUri(&articleIDForm); err != nil {
		panic(model.UnexpectedParamsError)
	}

	service.UpdateArticle(articleIDForm.ID, articleJsonForm)
	c.JSON(201, "操作成功")
}

func DeleteArticle(c *gin.Context) {
	articleUriForm := form.ArticleIDForm{}
	if err := c.ShouldBindUri(&articleUriForm); err != nil {
		panic(model.UnexpectedParamsError)
	}
	service.DeleteArticle(articleUriForm.ID)
	c.JSON(201, "操作成功")
}
