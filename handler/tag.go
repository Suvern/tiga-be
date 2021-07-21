package handler

import (
	"github.com/gin-gonic/gin"
	"tiga/model"
	"tiga/model/form"
	"tiga/service"
)

func GetTagList(c *gin.Context) {
	c.JSON(200, service.GetTagList())
}

func DeleteTag(c *gin.Context) {
	tagIDForm := form.TagIDForm{}
	if err := c.ShouldBindUri(tagIDForm); err != nil {
		panic(model.UnexpectedParamsError)
	}
	service.DeleteTag(tagIDForm.ID)
	c.JSON(201, "操作成功")
}
