package form

type ArticleForm struct {
	Title        string   `json:"title"`
	Content      string   `json:"content"`
	TagName      []string `json:"tagName"`
	CategoryName string   `json:"categoryName"`
}
