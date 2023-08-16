package config

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
)

func Get(key string) string {

	//获取文件根路径
	wd, _ := os.Getwd()
	// 设置配置文件的名字
	viper.SetConfigName("config")
	// 设置配置文件的类型
	viper.SetConfigType("yaml")
	// 添加配置文件的路径，指定 config 目录下寻找
	viper.AddConfigPath(wd)
	// 寻找配置文件并读取
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}

	//转换成string类型
	return fmt.Sprint(viper.Get(key))
}
