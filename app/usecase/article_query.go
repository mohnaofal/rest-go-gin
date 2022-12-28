package usecase

import (
	"context"
	"fmt"
	"strings"

	"github.com/mohnaofal/rest-go-gin/app/models"
	"github.com/mohnaofal/rest-go-gin/app/repository/queries"
	"github.com/mohnaofal/rest-go-gin/app/usecase/request"
	"github.com/rs/zerolog/log"
)

// articleQueryUsecase
type articleQueryUsecase struct {
	articleQuery queries.ArticleQuery
}

// ArticleQueryUsecase
type ArticleQueryUsecase interface {
	View(ctx context.Context, form *request.ViewArticle) ([]models.Article, error)
}

// NewArticleQueryUsecase
func NewArticleQueryUsecase(articleQuery queries.ArticleQuery) ArticleQueryUsecase {
	return &articleQueryUsecase{articleQuery: articleQuery}
}

func (c *articleQueryUsecase) View(ctx context.Context, form *request.ViewArticle) ([]models.Article, error) {

	params := &models.ArticleParams{
		Query: func() string {
			query := strings.TrimSpace(form.Query)
			if query != `` {
				return fmt.Sprintf(`%%%s%%`, form.Query)
			}
			return ``
		}(),
		Author:  form.Author,
		OrderBy: `id`,
		SortBy:  `DESC`,
	}

	data, err := c.articleQuery.Select(ctx, params)
	if err != nil {
		log.Err(err)
		return data, err
	}

	return data, nil
}
