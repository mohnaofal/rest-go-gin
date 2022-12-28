package commands

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/go-redis/redis"
	"github.com/mohnaofal/rest-go-gin/app/models"
	"github.com/mohnaofal/rest-go-gin/config"
	"github.com/mohnaofal/rest-go-gin/config/mysql"
	"github.com/rs/zerolog/log"
)

type articleCommand struct {
	cfg     *config.Config
	mysqlDB mysql.MySQLConnection
	cache   *redis.Client
}

type ArticleCommand interface {
	Insert(ctx context.Context, data *models.Article) (*models.Article, error)

	SetCache(ctx context.Context, key string, data interface{}) error
}

func NewArticleCommand(cfg *config.Config) ArticleCommand {
	return &articleCommand{
		cfg:     cfg,
		mysqlDB: cfg.MySQLDB(),
		cache:   cfg.Redis(),
	}
}

func (c *articleCommand) Insert(ctx context.Context, data *models.Article) (*models.Article, error) {
	commandExec, err := c.mysqlDB.MySQL().Exec(`INSERT INTO article (author, title, body, created) VALUES(?, ?, ?, ?)`, data.Author, data.Title, data.Body, data.Created)
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

	// Set Key
	key := fmt.Sprintf(`article:%v`, id)
	// Cache Article after insert
	c.SetCache(ctx, key, data)

	return data, nil
}

func (c *articleCommand) SetCache(ctx context.Context, key string, data interface{}) error {
	dataString, err := json.Marshal(data)
	if err != nil {
		return err
	}

	if err := c.cache.Set(key, string(dataString), time.Hour).Err(); err != nil {
		return err
	}
	return nil
}
