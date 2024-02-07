package controller

import (
	response "dpj-admin-api/support"
	"dpj-admin-api/support/http"
	client "dpj-admin-api/support/rabbit"
	"fmt"
	"github.com/gin-gonic/gin"
)

func GetNoticeList(c *gin.Context) {

	// http 客户端示例
	body, statusCode, err := http.NewHTTPClient().Get("https://www.baidu.com")
	if err != nil {
		fmt.Println("GET request failed:", err)
		return
	}
	fmt.Printf("GET Response Code: %d\n", statusCode)
	fmt.Printf("GET Response Body: %s\n", body)
	response.WithContext(c).Success("获取成功！")
}

func GetNotice(c *gin.Context) {

	// 创建 RabbitMQ 实例时将使用新的连接配置
	rabbitmq, _ := client.NewRabbitMQ("queueName")

	// 其他操作...
	rabbitmq.PublishSimple("Hello, RabbitMQ!")
	//rabbitmq.ConsumeSimple()
	// 最后别忘了关闭连接
	defer rabbitmq.Destroy()
	response.WithContext(c).Success("获取成功！")
}
