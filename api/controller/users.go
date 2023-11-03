package controller

import (
	"dpj-admin-api/config"
	response "dpj-admin-api/support"
	"github.com/gin-gonic/gin"
)

type DpjUser struct {
	ID            int `gorm:"primaryKey"`
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
type requestUserList struct {
	Page         int    `json:"page"`
	PageSize     int    `json:"size"`
	Nickname     string `json:"nickname"`
	Status       *int   `json:"status"`
	OnlineStatus *int   `json:"online_status"`
}

func UserList(c *gin.Context) {

	// 未处理page,size默认0
	var request requestUserList
	c.Bind(&request)

	// 开始执行查询
	query := config.Db().Model(&DpjUser{})

	if request.Nickname != "" {
		query.Where("nickname LIKE ?", "%"+request.Nickname+"%")
	}

	if request.Status != nil {
		query.Where("status = ?", request.Status)
	}

	if request.OnlineStatus != nil {
		query.Where("online_status = ?", request.OnlineStatus)
	}

	// 开始执行统计
	var userCount int64
	query.Count(&userCount)

	var DpjUserList []DpjUser
	query.Limit(request.PageSize).Offset((request.Page - 1) * request.PageSize).Find(&DpjUserList)

	response.WithContext(c).Success(gin.H{
		"total": userCount,
		"list":  DpjUserList,
	})
}
