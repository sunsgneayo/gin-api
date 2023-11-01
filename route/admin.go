package route

import (
	"dpj-admin-api/api/controller"
	"github.com/gin-gonic/gin"
)

func CommonRouter(r *gin.Engine) *gin.Engine {

	// 注册
	r.POST("register", controller.Register)

	// 登录
	r.POST("login", controller.Login)

	// 获取验证码
	r.GET("captcha", controller.Captcha)

	// 文件上传
	r.POST("upload", controller.UploadFile)

	return r
}
