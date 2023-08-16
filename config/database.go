package config

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Db() *gorm.DB {
	host := "43.138.132.9"
	port := "3398"
	username := "data_center"
	password := "KCMBfAjeJhbJXsSe"
	database := "data_center"
	charset := "utf8"
	args := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=true",
		username,
		password,
		host,
		port,
		database,
		charset)
	db, err := gorm.Open(mysql.Open(args))
	// 数据库连接出错 退出
	if err != nil {
		fmt.Println("数据库连接失败", err)
		panic(1)
	}

	// 返回数据库实例
	return db
}
