package handler

import "github.com/gin-gonic/gin"

func GetTagList(c *gin.Context) {
	c.JSON(200, "getTagList")
}

func CreateTag(c *gin.Context) {
	c.JSON(200, "createTag")
}

func DeleteTag(c *gin.Context) {
	c.JSON(200, "deleteTag")
}
