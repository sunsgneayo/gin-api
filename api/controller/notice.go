package controller

import (
	response "dpj-admin-api/support"
	"github.com/gin-gonic/gin"
)

func GetNoticeList(c *gin.Context) {

	response.WithContext(c).Success("获取成功！")
}

func GetNotice(c *gin.Context) {

	response.WithContext(c).Success("获取成功！")
}
