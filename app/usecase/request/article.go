package request

type ViewArticle struct {
	Query  string `form:"query" json:"query"`
	Author string `form:"author" json:"author"`
}
