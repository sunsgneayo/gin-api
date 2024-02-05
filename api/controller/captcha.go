package controller

import (
	"dpj-admin-api/api/service"
	response "dpj-admin-api/support"
	"github.com/gin-gonic/gin"
)

func CaptchaActive(c *gin.Context) {
	// Generate Captcha
	_, b64, tb64, key, err := service.GenerateCaptcha()
	if err != nil {
		panic(err)
		return
	}
	response.WithContext(c).Success(gin.H{
		"image_base64": b64,
		"thumb_base64": tb64,
		"captcha_key":  key,
		"code":         0,
	})

}
