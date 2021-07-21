package service

import (
	"tiga/db"
	"tiga/model/dao"
)

func GetTagList() []dao.Tag {
	var tagList []dao.Tag
	db.Db.Find(&tagList)
	return tagList
}

func DeleteTag(id uint) {
	db.Db.Delete(&dao.Tag{}, id)
}
