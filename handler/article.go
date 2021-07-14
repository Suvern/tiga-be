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
	createArticleForm := form.CreateArticleForm{}
	if err := c.ShouldBindJSON(&createArticleForm); err != nil {
		panic(model.UnexpectedParamsError)
	}
	service.CreateArticle(createArticleForm)
	c.JSON(201, "操作成功")
}

func UpdateArticle(c *gin.Context) {
	c.JSON(200, "updateArticle")
}

func DeleteArticle(c *gin.Context) {
	c.JSON(200, "deleteArticle")
}
