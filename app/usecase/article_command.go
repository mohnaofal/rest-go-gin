package usecase

import (
	"context"
	"time"

	"github.com/mohnaofal/rest-go-gin/app/models"
	"github.com/mohnaofal/rest-go-gin/app/repository/commands"
	"github.com/mohnaofal/rest-go-gin/app/repository/queries"
	"github.com/rs/zerolog/log"
)

// articleCommandUsecase
type articleCommandUsecase struct {
	articleCommand commands.ArticleCommand
	articleQuery   queries.ArticleQuery
}

type ArticleCommandUsecase interface {
	Create(ctx context.Context, form *models.Article) (*models.Article, error)
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

func (c *articleCommandUsecase) Create(ctx context.Context, form *models.Article) (*models.Article, error) {
	form.Created = time.Now()
	form, err := c.articleCommand.Insert(ctx, form)
	if err != nil {
		log.Err(err)
		return nil, err
	}

	return form, nil
}
