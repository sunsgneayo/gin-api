package route

import (
	"dpj-admin-api/api/controller"
	"github.com/gin-gonic/gin"
)

func CommonRouter(r *gin.Engine) *gin.Engine {

	r.POST("register", controller.Register)
	r.POST("login", controller.Login)
	r.GET("captcha", controller.Captcha)

	r.POST("upload", controller.UploadFile)

	return r
}
