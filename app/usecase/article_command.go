package usecase

import (
	"context"
	"fmt"
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
	now, _ := time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	form.Created = now
	data, err := c.articleCommand.Insert(ctx, form)
	if err != nil {
		log.Err(err)
		return nil, err
	}

	// Set Key
	key := fmt.Sprintf(`article:%v`, data.ID)
	// Cache Article
	c.articleCommand.SetCache(ctx, key, data, time.Hour)

	return form, nil
}
