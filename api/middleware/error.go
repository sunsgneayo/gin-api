package middleware

import (
	response "dpj-admin-api/support"
	"github.com/gin-gonic/gin"
	"log"
	"runtime/debug"
)

// ErrorHandle  Error 统一500错误处理函数
func ErrorHandle(c *gin.Context) {
	defer func() {
		if r := recover(); r != nil {
			//打印错误堆栈信息
			log.Printf("panic: %v\n", r)
			debug.PrintStack()
			//封装通用json返回
			response.WithContext(c).Error(500, "服务异常")
		}
	}()
	c.Next()
}
