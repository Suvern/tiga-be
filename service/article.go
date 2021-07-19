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

func GetArticleDetail(id uint) dao.Article {
	if !isArticleExist(id) {
		panic(model.HttpError{Code: http.StatusNotFound, Msg: "文章不存在或已删除"})
	}
	article := dao.Article{}
	db.Db.Where(id).Find(&article)
	return article
}

func CreateArticle(form form.ArticleJsonForm) {
	article := dao.Article{
		Title:   form.Title,
		Content: form.Content,
	}
	setTagAndCategoryWithForm(&article, form)
	db.Db.Save(&article)
}

func UpdateArticle(id uint, form form.ArticleJsonForm) {
	if !isArticleExist(id) {
		panic(model.HttpError{Code: http.StatusNotFound, Msg: "文章不存在或已删除"})
	}
	article := dao.Article{}
	db.Db.Where(id).Find(&article)
	article.Title = form.Title
	article.Content = form.Content
	setTagAndCategoryWithForm(&article, form)
	db.Db.Save(&article)
}

func DeleteArticle(id uint) {
	if !isArticleExist(id) {
		panic(model.HttpError{Code: http.StatusNotFound, Msg: "文章不存在或已删除"})
	}
	db.Db.Where(id).Delete(&dao.Article{})
}

func isArticleExist(id uint) bool {
	article := dao.Article{}
	result := db.Db.Where(id).Find(&article)
	return !errors.Is(result.Error, gorm.ErrRecordNotFound)
}

func setTagAndCategoryWithForm(article *dao.Article, form form.ArticleJsonForm) {
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
