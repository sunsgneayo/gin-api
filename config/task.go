package config

import (
	"fmt"
	"github.com/robfig/cron/v3"
)

// newWithSeconds  返回一个支持至 秒 级别的 cron
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
	list := [5]int{1, 2, 3, 4, 5}
	for _, attribute := range list {
		// 如果不给attribute重新赋值引用那么Println永远是数组的最后一个值
		attribute := attribute
		t.AddFunc(spec, func() {
			fmt.Println("cron running:", attribute)
		})
	}

	// 任务开始
	t.Start()

	select {}
}
