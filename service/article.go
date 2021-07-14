package service

import (
	"tiga/db"
	"tiga/model/dao"
	"tiga/model/form"
)

func GetArticleList() {

}

func CreateArticle(form form.CreateArticleForm) {
	article := dao.Article{
		Title:   form.Title,
		Content: form.Content,
	}
	if len(form.TagName) > 0 && len(form.TagName[0]) > 0 {
		var tags []dao.Tag
		db.Db.Model(&dao.Tag{}).Where("name IN ?", form.TagName).Find(&tags)
		article.Tags = tags
	}
	if len(form.CategoryName) > 0 {
		category := dao.Category{}
		db.Db.Model(&dao.Category{}).Where("name = ?", form.CategoryName).Find(&category)
		article.CategoryID = category.ID
	}
	db.Db.Save(&article)
}
