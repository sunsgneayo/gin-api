package route

import (
	"dpj-admin-api/api/controller"
	"dpj-admin-api/api/middleware"
	"github.com/gin-gonic/gin"
)

func AdminRouter(r *gin.RouterGroup) {
	router := r.Group("admin").Use(middleware.JwtAuth())

	// 获取管理员信息
	router.POST("info", controller.UserInfo)

}
