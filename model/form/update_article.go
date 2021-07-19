package form

type ArticleUriForm struct {
	ID uint `uri:"id" binding:"required"`
}
