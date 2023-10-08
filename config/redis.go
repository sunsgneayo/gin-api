package config

import (
	"context"
	"github.com/go-redis/redis/v8"
)

var (
	RedisDb *redis.Client
)

// SetupRedisDb

func SetupRedisDb() error {
	var ctx = context.Background()
	RedisDb = redis.NewClient(&redis.Options{
		//Addr:     Get("redis.host") + ":" + Get("redis.port"),
		Addr:     "43.138.132.9:6379",
		Password: "fanjiao2021", // no password set
		DB:       0,             // use default DB
	})

	_, err := RedisDb.Ping(ctx).Result()
	if err != nil {
		return err
	}
	return nil
}
