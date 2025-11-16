package common

import (
	"context"
	"os"

	"github.com/redis/go-redis/v9"
)

var Redis *redis.Client
var Ctx = context.Background()

func InitRedis() {
	Redis = redis.NewClient(&redis.Options{
		Addr: os.Getenv("REDIS_URL"),
		Password: "",
		DB: 0,
	})

	if err := Redis.Ping(Ctx).Err(); err != nil {
        panic("Failed to connect to Redis: " + err.Error())
    }
}