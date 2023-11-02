package controller

import (
	"dpj-admin-api/config"
	response "dpj-admin-api/support"
	"github.com/gin-gonic/gin"
)

type requestUserList struct {
	Username string `json:"username"`
}

type DpjUser struct {
	ID            int
	OpenId        string
	Nickname      string
	HeadPicture   string
	OnlineStatus  int
	Status        int
	Balance       int
	CreateTime    string
	LastLoginTime string
	InRoomId      int
	MatchNumber   int
	Region        string
}

func UserList(c *gin.Context) {

	//var request requestUserList
	//err := c.Bind(&request)
	//if err != nil {
	//	response.WithContext(c).Error(400, "参数获取失败")
	//	return
	//}

	var userCount int64
	config.Db().Model(&DpjUser{}).Count(&userCount)

	var DpjUserList []DpjUser
	config.Db().Model(&DpjUser{}).Limit(10).Find(&DpjUserList)

	response.WithContext(c).Success(gin.H{
		"total": userCount,
		"list":  DpjUserList,
	})
}
