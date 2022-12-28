package redis

import (
	"fmt"
	"os"
	"strconv"

	"github.com/go-redis/redis"
)

func InitRedis() *redis.Client {
	db, _ := strconv.Atoi(os.Getenv("REDIS_DB"))
	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%v:%v", os.Getenv("REDIS_HOST"), os.Getenv("REDIS_PORT")),
		Password: os.Getenv("REDIS_PASSWORD"),
		DB:       db,
	})

	if err := rdb.Ping().Err(); err != nil {
		panic(err)
	}

	return rdb
}
