package initialize

import "dpj-admin-api/config"

func TaskServer() {
	// 初始化 异步任务调度服务
	if config.Get("app.corn_task") != "true" {
		go func() {
			config.InitTaskRun()
		}()
	}
}
