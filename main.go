package main

import (
	"dpj-admin-api/config"
	"dpj-admin-api/route"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.New()
	r.Use(gin.Logger())

	// 连接redis
	config.SetupRedisDb()

	route.InitRoute(r)

	port := config.Get("app.default_listen_port")
	host := config.Get("app.default_local_host")

	// 服务初始化
	r.Run(host + ":" + port)

	// 定时任务初始化（需要单独运行）
	config.InitTaskRun()
}
