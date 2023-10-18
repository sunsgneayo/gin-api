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
		//Addr:     Get("redis.host") + ":" + Get("redis.port"),
		Addr:     "43.138.132.9:6379",
		Password: "fanjiao2022", // no password set
		DB:       1,             // use default DB
	})

	_, err := RedisDb.Ping(ctx).Result()
	if err != nil {
		return err
	}
	return nil
}
