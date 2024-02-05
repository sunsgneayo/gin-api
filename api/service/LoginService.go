package service

import (
	"dpj-admin-api/api/middleware"
	"dpj-admin-api/config"
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

type DpjAdmins struct {
	Username string
	Password string
	ID       int
}

// Login /** 行为验证码登录服务
func Login(username string, password string) (string, error) {

	var user DpjAdmins
	config.Db().Where("username = ?", username).First(&user)
	if user.ID == 0 {
		return "", fmt.Errorf("用户不存在")
	}

	//判断密码是否正确
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return "", fmt.Errorf("密码错误")
	}

	tokenString, _ := middleware.GenToken(user.Username, user.ID)
	return tokenString, nil
}
