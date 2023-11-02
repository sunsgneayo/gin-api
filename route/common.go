package route

import (
	"dpj-admin-api/api/controller"
	"github.com/gin-gonic/gin"
)

func LoginRouter(r *gin.RouterGroup) {

	// 注册
	r.POST("register", controller.Register)

	// 登录
	r.POST("login", controller.Login)

	// 退出登录
	r.POST("login_out", controller.LoginOut)

	// 获取验证码
	r.GET("captcha", controller.Captcha)

	// 文件上传
	r.POST("upload", controller.UploadFile)

}
