package route

import "github.com/gin-gonic/gin"

func InitRoute(r *gin.Engine) {

	// 默认加载log日志中间件
	r.Use(gin.Logger())

	// 默认使用api分组
	api := r.Group("/api")

	// 开始分配路由

	LoginRouter(api)

	NoticeRouter(api)
}
