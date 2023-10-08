package route

import (
	"dpj-admin-api/http/controller"
	"github.com/gin-gonic/gin"
)

func InitRouter(r *gin.Engine) *gin.Engine {

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.POST("register", controller.Register)
	r.POST("login", controller.Login)

	r.GET("captcha", controller.Captcha)

	return r
}
