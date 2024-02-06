package initialize

import (
	"dpj-admin-api/config"
	"log"
)

func RedisServer() {
	// 链接redis 服务
	if config.Get("redis.host") != "" {
		// 连接redis
		err := config.SetupRedisDb()
		if err != nil {
			log.Println("连接redis失败...", err)
			return
		}
	}

}
