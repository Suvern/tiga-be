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

type ArticleVO struct {
	Title        string
	Tags         []dao.Tag `gorm:"many2many:article_tag"`
	CategoryType string
	CategoryID   uint
}

const PageSize = 15

func GetArticleList(form form.ArticleQueryForm) []ArticleVO {
	var articles []ArticleVO
	result := db.Db.Model(&dao.Article{})
	if form.TagID != nil {
		result = result.Joins("Tags").Joins("tag.name = ?", form.TagID)
	}
	if form.CategoryID != nil {
		result = result.Joins("Category").Joins("category.name = ?", form.TagID)
	}
	result.Limit(PageSize).Offset(form.Page * PageSize).Find(&articles)
	return articles
}

func GetArticleDetail(id uint) dao.Article {
	if !isArticleExist(id) {
		panic(model.HttpError{Code: http.StatusNotFound, Msg: "文章不存在或已删除"})
	}
	var article dao.Article
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
	var article dao.Article
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
	var article dao.Article
	result := db.Db.Where(id).Find(&article)
	return !errors.Is(result.Error, gorm.ErrRecordNotFound)
}

func setTagAndCategoryWithForm(article *dao.Article, form form.ArticleJsonForm) {
	if len(form.TagID) > 0 && len(form.TagID[0]) > 0 {
		var tags []dao.Tag
		db.Db.Model(&dao.Tag{}).Where("name IN ?", form.TagID).Find(&tags)
		article.Tags = tags
	}
	if len(form.CategoryID) > 0 {
		var category dao.Category
		db.Db.Model(&dao.Category{}).Where("name = ?", form.CategoryID).Find(&category)
		article.CategoryID = category.ID
	}
}
