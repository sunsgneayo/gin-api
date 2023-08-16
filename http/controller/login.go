package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/mojocn/base64Captcha"
)

func login(c *gin.Engine) {

}

type User struct {
	UserName string
	Password string
}

func Register(c *gin.Context) {

	if store.Verify(c.PostForm("captchaId"), c.PostForm("captcha"), true) == false {
		c.JSON(400, gin.H{
			"message": "验证码错误",
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "ok",
		//"res":     result.RowsAffected,
		"ress": c.BindJSON("a"),
	})
	//user := User{UserName: c.Param("username"), Password: c.Param("password")}
	//
	//result := config.Db().Create(&user) // 通过数据的指针来创建
	//
	//if result.Error != nil {
	//	c.JSON(400, gin.H{
	//		"message": result.Error,
	//	})
	//	return
	//}

}

type configJsonBody struct {
	Id          string
	CaptchaType string
	VerifyValue string
	DriverMath  *base64Captcha.DriverMath
}

var store = base64Captcha.DefaultMemStore

func Captcha(c *gin.Context) {
	//定义一个driver
	var driver base64Captcha.Driver
	//创建一个字符串类型的验证码驱动DriverString, DriverMath :算式驱动
	driverString := base64Captcha.DriverMath{
		Height:          40,    //高度
		Width:           100,   //宽度
		NoiseCount:      0,     //干扰数
		ShowLineOptions: 3 | 4, //展示个数
	}
	driver = driverString.ConvertFonts()
	//生成验证码
	cap := base64Captcha.NewCaptcha(driver, store)
	id, b64s, err := cap.Generate()

	if err != nil {

	}
	c.JSON(200, gin.H{
		"message":   "ok",
		"captcha":   b64s,
		"captchaId": id,
	})
}
