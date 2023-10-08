package controller

import (
	"dpj-admin-api/config"
	response "dpj-admin-api/support"
	"github.com/gin-gonic/gin"
	"github.com/mojocn/base64Captcha"
	"golang.org/x/crypto/bcrypt"
)

func Login(c *gin.Context) {

	if store.Verify(c.PostForm("captchaId"), c.PostForm("captcha"), true) == false {
		response.WithContext(c).Error(400, "验证码错误")
		return
	}

	//获取参数
	username := c.PostForm("username")
	password := c.PostForm("password")

	//数据验证
	if len(username) < 4 {
		response.WithContext(c).Error(400, "用户名不能小于4位")
		return
	}
	if len(password) < 6 {
		response.WithContext(c).Error(400, "密码不能小于6位")
		return
	}

	//判断手机号是否存在
	var user DpjAdmin
	config.Db().Where("username = ?", username).First(&user)
	if user.ID == 0 {
		response.WithContext(c).Error(400, "用户不存在")
		return
	}

	//判断密码是否正确
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		response.WithContext(c).Error(400, "密码错误")
		return
	}

	response.WithContext(c).Success("登录成功")

}

type DpjAdmin struct {
	Username string
	Password string
	ID       int
}

func Register(c *gin.Context) {

	println(c.PostForm("captchaId"))
	println(c.PostForm("captcha"))

	an := store.Get(c.PostForm("captchaId"), false)

	println(an)

	if store.Verify(c.PostForm("captchaId"), c.PostForm("captcha"), false) == false {
		response.WithContext(c).Error(400, "验证码错误")
		return
	}

	//获取参数
	username := c.PostForm("username")
	password := c.PostForm("password")

	//数据验证
	if len(username) == 0 {
		response.WithContext(c).Error(400, "用户名不能为空")
		return
	}
	if len(password) < 6 {
		response.WithContext(c).Error(400, "密码不能少于6位")
		return
	}

	//判断手机号是否存在
	var user DpjAdmin
	config.Db().Where("username = ?", username).First(&user)
	if user.ID != 0 {
		response.WithContext(c).Error(400, "用户已存在")
		return
	}

	//创建用户
	Password, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		response.WithContext(c).Error(400, "密码加密错误")
		return
	}
	newUser := DpjAdmin{
		Username: username,
		Password: string(Password),
	}
	config.Db().Create(&newUser)

	response.WithContext(c).Success("注册成功！")
}

var store = base64Captcha.DefaultMemStore

// Captcha 获取验证码
func Captcha(c *gin.Context) {
	//定义一个driver
	var driver base64Captcha.Driver
	//创建一个字符串类型的验证码驱动DriverString, DriverMath :算式驱动
	driverString := base64Captcha.DriverMath{
		Height:          50,    //高度
		Width:           120,   //宽度
		NoiseCount:      0,     //干扰数
		ShowLineOptions: 3 | 4, //展示个数
	}
	driver = driverString.ConvertFonts()
	//生成验证码
	cap := base64Captcha.NewCaptcha(driver, store)
	id, b64s, err := cap.Generate()
	store.Set(id, b64s)
	if err != nil {
		response.WithContext(c).Error(500, "Server Error")
		return
	}
	response.WithContext(c).Success(gin.H{
		"captcha":   b64s,
		"captchaId": id,
	})
}
