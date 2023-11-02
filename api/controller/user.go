package controller

import (
	response "dpj-admin-api/support"
	"github.com/gin-gonic/gin"
)

func UserInfo(c *gin.Context) {

	response.WithContext(c).Success(gin.H{
		"name":   "王立群",
		"avatar": "",
	})

}
