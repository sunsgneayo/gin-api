package route

import (
	"dpj-admin-api/api/middleware"
	"github.com/gin-gonic/gin"
)

func InitRoute(r *gin.Engine) {

	// 默认加载log日志中间件
	r.Use(gin.Logger())
	// 加载跨域中间件
	r.Use(middleware.Cors())

	// 默认使用api分组
	api := r.Group("/api")

	// 开始分配路由

	LoginRouter(api)

	NoticeRouter(api)
}
