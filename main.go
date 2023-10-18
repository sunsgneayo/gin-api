package main

import (
	"dpj-admin-api/config"
	"dpj-admin-api/route"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.New()
	r.Use(gin.Logger())
	//r.Use(middleware.Casbin())

	config.SetupRedisDb()
	route.InitRouter(r)

	port := config.Get("app.default_listen_port")
	host := config.Get("app.default_local_host")

	r.Run(host + ":" + port) // 监听并在 0.0.0.0:8008 上启动服务
}
