package infrastructure

import (
	"fmt"
	"os"

	"github.com/redis/go-redis/v9"
)

var (
	RedisClient *redis.Client
)

func GetRedisClient() error {
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", os.Getenv("REDIS_HOST"), os.Getenv("REDIS_PORT")),
		Password: os.Getenv("REDIS_PASSWORD"),
		DB:       0,
	})

	return nil
}
