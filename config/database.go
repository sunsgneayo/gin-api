package config

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Db() *gorm.DB {
	host := Get("mysql.host")
	port := Get("mysql.port")
	username := Get("mysql.username")
	password := Get("mysql.password")
	database := Get("mysql.database")
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
