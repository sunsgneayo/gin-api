package controller

import (
	response "dpj-admin-api/support"
	"dpj-admin-api/support/client"
	"fmt"
	"github.com/gin-gonic/gin"
)

func GetNoticeList(c *gin.Context) {

	// http 客户端示例
	body, statusCode, err := client.NewHTTPClient().Get("https://www.baidu.com")
	if err != nil {
		fmt.Println("GET request failed:", err)
		return
	}
	fmt.Printf("GET Response Code: %d\n", statusCode)
	fmt.Printf("GET Response Body: %s\n", body)
	response.WithContext(c).Success("获取成功！")
}

func GetNotice(c *gin.Context) {

	response.WithContext(c).Success("获取成功！")
}
