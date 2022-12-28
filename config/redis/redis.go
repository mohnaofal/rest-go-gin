package redis

import (
	"os"

	"github.com/go-redis/redis"
)

func InitRedis() *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_HOST"),
		Password: os.Getenv("REDIS_PASSWORD"), // no password set
		DB:       0,                           // use default DB
	})

	if err := rdb.Ping().Err(); err != nil {
		panic(err)
	}

	return rdb
}
