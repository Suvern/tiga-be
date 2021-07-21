package form

type ArticleJsonForm struct {
	Title      string   `json:"title" binding:"required"`
	Content    string   `json:"content" binding:"required"`
	TagID      []string `json:"tagName" binding:"required"`
	CategoryID uint     `json:"categoryName" binding:"required"`
}

type ArticleIDForm struct {
	ID uint `uri:"id" binding:"required"`
}

type ArticleQueryForm struct {
	Page       int     `form:"page,omitempty"`
	CategoryID *string `form:"category,omitempty"`
	TagID      *string `form:"tag,omitempty"`
}
