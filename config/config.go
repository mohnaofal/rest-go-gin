package config

import (
	"github.com/go-redis/redis"
	"github.com/mohnaofal/rest-go-gin/config/mysql"
	rdb "github.com/mohnaofal/rest-go-gin/config/redis"
)

type Config struct {
	mysqlDB mysql.MySQLConnection
	redis   *redis.Client
}

func LoadConfig() *Config {
	cfg := new(Config)

	cfg.InitMySQLDB() // mysql
	cfg.InitRedis()   // redis

	return cfg
}

func (c *Config) InitMySQLDB() {
	c.mysqlDB = mysql.ConnectionDB()
}

func (c *Config) InitRedis() {
	c.redis = rdb.InitRedis()
}

func (c *Config) MySQLDB() mysql.MySQLConnection {
	return c.mysqlDB
}

func (c *Config) Redis() *redis.Client {
	return c.redis
}
