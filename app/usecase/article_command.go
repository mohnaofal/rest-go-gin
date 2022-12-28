package usecase

import (
	"github.com/mohnaofal/rest-go-gin/app/repository/commands"
	"github.com/mohnaofal/rest-go-gin/app/repository/queries"
)

// articleCommandUsecase
type articleCommandUsecase struct {
	articleCommand commands.ArticleCommand
	articleQuery   queries.ArticleQuery
}

type ArticleCommandUsecase interface {
}

func NewArticleCommandUsecase(
	articleCommand commands.ArticleCommand,
	articleQuery queries.ArticleQuery,
) ArticleCommandUsecase {
	return &articleCommandUsecase{
		articleCommand: articleCommand,
		articleQuery:   articleQuery,
	}
}
