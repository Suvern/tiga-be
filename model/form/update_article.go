package form

type UpdateArticleForm struct {
	ID uint `uri:"id" binding:"required"`
}
