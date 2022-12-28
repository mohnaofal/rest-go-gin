package usecase

import "github.com/mohnaofal/rest-go-gin/app/repository/queries"

// articleQueryUsecase
type articleQueryUsecase struct {
	articleQuery queries.ArticleQuery
}

// ArticleQueryUsecase
type ArticleQueryUsecase interface {
}

// NewArticleQueryUsecase
func NewArticleQueryUsecase(articleQuery queries.ArticleQuery) ArticleQueryUsecase {
	return &articleQueryUsecase{articleQuery: articleQuery}
}
