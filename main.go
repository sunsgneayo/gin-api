package main

import "dpj-admin-api/initialize"

func main() {
	// 异步任务调度服务
	initialize.TaskServer()
	// 初始化redis连接
	initialize.RedisServer()
	// 队列消费服务
	initialize.RabbitConsume()

	// HTTP服务 -> 放在最后，否则在它之后的服务都不能正常运行
	initialize.HttpMainServer()

}
