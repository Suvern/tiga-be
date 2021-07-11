package handler

import "github.com/gin-gonic/gin"

func GetArticleList(c *gin.Context) {
	c.JSON(200, "getArticleList")
}

func GetArticleDetail(c *gin.Context) {
	c.JSON(200, "getArticle")
}

func CreateArticle(c *gin.Context) {
	c.JSON(200, "createArticle")
}

func UpdateArticle(c *gin.Context) {
	c.JSON(200, "updateArticle")
}

func DeleteArticle(c *gin.Context) {
	c.JSON(200, "deleteArticle")
}
