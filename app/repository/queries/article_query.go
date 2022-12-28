package queries

import (
	"context"
	"fmt"
	"time"

	"github.com/mohnaofal/rest-go-gin/app/models"
	"github.com/mohnaofal/rest-go-gin/config"
	"github.com/mohnaofal/rest-go-gin/config/mysql"
	"github.com/rs/zerolog/log"
)

type articleQuery struct {
	cfg     *config.Config
	mysqlDB mysql.MySQLConnection
}

type ArticleQuery interface {
	Select(ctx context.Context, params *models.ArticleParams) ([]models.Article, error)
}

func NewArticleQuery(cfg *config.Config) ArticleQuery {
	return &articleQuery{
		cfg:     cfg,
		mysqlDB: cfg.MySQLDB(),
	}
}

func (c *articleQuery) Select(ctx context.Context, params *models.ArticleParams) ([]models.Article, error) {
	var (
		res            = make([]models.Article, 0)
		queryCondition = ``
	)

	if params.Query != `` {
		queryCondition += func() string {
			if len(queryCondition) > 0 {
				return fmt.Sprintf(` AND (title LIKE "%s" OR body LIKE "%s)"`, params.Query, params.Query)
			}
			return fmt.Sprintf(`WHERE (title LIKE "%s" OR body LIKE "%s")`, params.Query, params.Query)
		}()
	}

	if params.Author != `` {
		queryCondition += func() string {
			if len(queryCondition) > 0 {
				return fmt.Sprintf(` AND author = "%s"`, params.Author)
			}
			return fmt.Sprintf(`WHERE author = "%s"`, params.Author)
		}()
	}

	if params.OrderBy != `` && params.SortBy != `` {
		queryCondition += fmt.Sprintf(` ORDER BY %s %s`, params.OrderBy, params.SortBy)
	}

	sqlPrepare, err := c.mysqlDB.MySQL().Prepare(`SELECT * FROM article ` + queryCondition)
	if err != nil {
		log.Err(err)
		return res, err
	}

	rows, err := sqlPrepare.Query()
	if err != nil {
		log.Err(err)
		return res, err
	}

	for rows.Next() {
		var (
			id                  uint
			author, title, body string
			created             time.Time
		)

		err = rows.Scan(&id, &author, &title, &body, &created)
		if err != nil {
			log.Err(err)
			return res, err
		}

		res = append(res, models.Article{
			ID:      id,
			Author:  author,
			Title:   title,
			Body:    body,
			Created: created,
		})
	}

	return res, nil
}
