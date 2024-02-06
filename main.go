package main

import "dpj-admin-api/initialize"

func main() {
	// 异步任务调度服务
	initialize.TaskServer()
	// 初始化redis连接
	initialize.RedisServer()
	// HTTP服务
	initialize.HttpMainServer()

}
