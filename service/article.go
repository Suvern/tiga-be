package service

import (
	"errors"
	"gorm.io/gorm"
	"net/http"
	"tiga/db"
	"tiga/model"
	"tiga/model/dao"
	"tiga/model/form"
)

func GetArticleList() {

}

func CreateArticle(form form.ArticleForm) {
	article := dao.Article{
		Title:   form.Title,
		Content: form.Content,
	}
	setTagAndCategoryWithForm(&article, form)
	db.Db.Save(&article)
}

func UpdateArticle(id uint, form form.ArticleForm) {
	article := dao.Article{}
	result := db.Db.Where(id).Find(&article)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		panic(model.HttpError{Code: http.StatusNotFound, Msg: "文章不存在或已删除"})
	}
	article.Title = form.Title
	article.Content = form.Content
	setTagAndCategoryWithForm(&article, form)
	db.Db.Save(&article)
}

func setTagAndCategoryWithForm(article *dao.Article, form form.ArticleForm) {
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
}
