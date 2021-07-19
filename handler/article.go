package handler

import (
	"github.com/gin-gonic/gin"
	"tiga/model"
	"tiga/model/form"
	"tiga/service"
)

func GetArticleList(c *gin.Context) {
	c.JSON(200, "getArticleList")
}

func GetArticleDetail(c *gin.Context) {
	c.JSON(200, "getArticle")
}

func CreateArticle(c *gin.Context) {
	articleForm := form.ArticleForm{}
	if err := c.ShouldBindJSON(&articleForm); err != nil {
		panic(model.UnexpectedParamsError)
	}
	service.CreateArticle(articleForm)
	c.JSON(201, "操作成功")
}

func UpdateArticle(c *gin.Context) {
	articleForm := form.ArticleForm{}
	if err := c.ShouldBindJSON(&articleForm); err != nil {
		panic(model.UnexpectedParamsError)
	}

	updateArticleForm := form.UpdateArticleForm{}
	if err := c.ShouldBindUri(&updateArticleForm); err != nil {
		panic(model.UnexpectedParamsError)
	}

	service.UpdateArticle(updateArticleForm.ID, articleForm)
	c.JSON(201, "操作成功")
}

func DeleteArticle(c *gin.Context) {
	c.JSON(200, "deleteArticle")
}
