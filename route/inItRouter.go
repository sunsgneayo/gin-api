package route

import (
	"dpj-admin-api/api/middleware"
	response "dpj-admin-api/support"
	"github.com/gin-gonic/gin"
)

func InitRoute(r *gin.Engine) {

	// 默认加载log日志中间件
	r.Use(gin.Logger())
	// 加载全部异常处理
	r.Use(middleware.ErrorHandle)
	// 加载跨域中间件
	r.Use(middleware.Cors())

	// 初始化404/405 异常结构
	r.NoRoute(func(c *gin.Context) {
		response.WithContext(c).Error(404, "not found router")
	})
	r.NoMethod(func(c *gin.Context) {
		response.WithContext(c).Error(405, "method not  found ")
	})
	// 默认使用api分组
	api := r.Group("/api")

	// 开始分配路由

	// 登录相关路由
	LoginRouter(api)

	NoticeRouter(api)

	// 管理员相关路由
	AdminRouter(api)

	// 平台用户列表关系路由
	UsersRouter(api)

	// 平台房间信息路由
	RoomsRouter(api)
}
