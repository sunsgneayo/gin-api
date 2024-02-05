package route

import (
	"dpj-admin-api/api/controller"
	"github.com/gin-gonic/gin"
)

func LoginRouter(r *gin.RouterGroup) {

	// 注册
	r.POST("register", controller.Register)

	// 登录
	r.POST("login", controller.LoginWithCaptcha)

	// 退出登录
	r.POST("login_out", controller.LoginOut)

	// 获取验证码
	r.GET("captcha", controller.Captcha)

	// 获取行为验证码
	r.GET("get_captcha", controller.CaptchaActive)

	//r.POST("check_captcha", controller.CheckCaptcha)

	// 文件上传
	r.POST("upload", controller.UploadFile)

}
