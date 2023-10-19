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

	//fmt.Println(ctx)
	var ctx = context.Background()
	RedisDb = redis.NewClient(&redis.Options{
		Addr:     Get("redis.host") + ":" + Get("redis.port"),
		Password: Get("redis.password"),
		DB:       0, // use default DB
	})

	_, err := RedisDb.Ping(ctx).Result()
	if err != nil {
		return err
	}
	return nil
}
