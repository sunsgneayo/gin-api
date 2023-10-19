package config

import (
	"dpj-admin-api/task"
	"fmt"
	"github.com/robfig/cron/v3"
)

//newWithSeconds  返回一个支持至 秒 级别的 cron
func newWithSeconds() *cron.Cron {
	secondParser := cron.NewParser(cron.Second | cron.Minute |
		cron.Hour | cron.Dom | cron.Month | cron.DowOptional | cron.Descriptor)
	return cron.New(cron.WithParser(secondParser), cron.WithChain())
}

// InitTaskRun 初始化定时任务
func InitTaskRun() {
	t := newWithSeconds()
	// corn 表达式
	spec := "0 */1 * * * ?"

	// 任务追加
	t.AddFunc(spec, task.TestABCBCB)

	i := 0
	t.AddFunc(spec, func() {
		i++
		fmt.Println("cron running:", i)
	})

	// 任务开始
	t.Start()

	select {}
}
