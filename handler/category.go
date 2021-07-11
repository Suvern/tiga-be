package handler

import "github.com/gin-gonic/gin"

func GetCategoryList(c *gin.Context) {
	c.JSON(200, "getCategoryList")
}

func CreateCategory(c *gin.Context) {
	c.JSON(200, "createCategoryList")
}

func UpdateCategory(c *gin.Context) {
	c.JSON(200, "updateCategory")
}

func DeleteCategory(c *gin.Context) {
	c.JSON(200, "deleteCategory")
}
