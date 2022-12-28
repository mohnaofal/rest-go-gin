package queries

import (
	"context"
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
	rows := make([]models.Article, 0)
	queryExec, err := c.mysqlDB.MySQL().Query(`SELECT * FROM article WHERE title LIKE ? OR body LIKE ? AND author = ?`, params.Query, params.Query, params.Author)
	if err != nil {
		log.Err(err)
		return rows, err
	}

	for queryExec.Next() {
		var (
			id                  uint
			author, title, body string
			created             time.Time
		)

		err = queryExec.Scan(&id, &author, &title, &body, &created)
		if err != nil {
			log.Err(err)
			return rows, err
		}

		rows = append(rows, models.Article{
			ID:      id,
			Author:  author,
			Title:   title,
			Body:    body,
			Created: created,
		})
	}

	return rows, nil
}
