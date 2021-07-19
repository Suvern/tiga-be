package form

type ArticleJsonForm struct {
	Title        string   `json:"title" binding:"required"`
	Content      string   `json:"content" binding:"required"`
	TagName      []string `json:"tagName" binding:"required"`
	CategoryName string   `json:"categoryName" binding:"required"`
}

type ArticleUriForm struct {
	ID uint `uri:"id" binding:"required"`
}

type ArticleQueryForm struct {
	CategoryName *string `form:"category,omitempty"`
	TagName      *string `form:"tag,omitempty"`
	Page         int     `form:"page,omitempty"`
}
