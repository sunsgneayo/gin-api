package controller

import (
	"dpj-admin-api/config"
	response "dpj-admin-api/support"
	"github.com/gin-gonic/gin"
	"net/http"
)

type requestRoomList struct {
	Page     int `json:"page"`
	PageSize int `json:"size"`
}

type DpjRoom struct {
	ID                       int
	RoomNumber               int
	ActivityId               string
	ActivityIdExpirationTime int
	MasterUserID             uint `gorm:"column:master_user_id;comment:'房主ID'"`
	CreateTime               string
	DissolveTime             string
	ShareQr                  string
	User                     DpjUser `gorm:"foreignKey:id;references:master_user_id"`
}

func RoomList(c *gin.Context) {

	var request requestRoomList
	err := c.Bind(&request)
	if err != nil {
		response.WithContext(c).Error(http.StatusBadRequest, "参数获取失败")
		return
	}

	var count int64
	config.Db().Model(&DpjRoom{}).Count(&count)

	var DpjRoomList []DpjRoom
	config.Db().Model(&DpjRoom{}).Preload("User").Limit(request.PageSize).Offset((request.Page - 1) * request.PageSize).Find(&DpjRoomList)

	response.WithContext(c).Success(gin.H{
		"total": count,
		"list":  DpjRoomList,
	})
}
