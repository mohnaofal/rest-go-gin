package commands

import (
	"context"

	"github.com/mohnaofal/rest-go-gin/app/models"
	"github.com/mohnaofal/rest-go-gin/config"
	"github.com/mohnaofal/rest-go-gin/config/mysql"
	"github.com/rs/zerolog/log"
)

type articleCommand struct {
	cfg     *config.Config
	mysqlDB mysql.MySQLConnection
}

type ArticleCommand interface {
	Insert(ctx context.Context, data *models.Article) (*models.Article, error)
}

func NewArticleCommand(cfg *config.Config) ArticleCommand {
	return &articleCommand{
		cfg:     cfg,
		mysqlDB: cfg.MySQLDB(),
	}
}

func (c *articleCommand) Insert(ctx context.Context, data *models.Article) (*models.Article, error) {
	commandExec, err := c.mysqlDB.MySQL().Exec(`INSERT INTO article(author, title, body) VALUES(?, ?, ?)`)
	if err != nil {
		log.Err(err)
		return nil, err
	}

	id, err := commandExec.LastInsertId()
	if err != nil {
		log.Err(err)
		return nil, err
	}

	data.ID = uint(id)

	return data, nil
}
