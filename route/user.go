package route

import (
	"dpj-admin-api/api/controller"
	"dpj-admin-api/api/middleware"
	"github.com/gin-gonic/gin"
)

func UsersRouter(r *gin.RouterGroup) {
	router := r.Group("user").Use(middleware.Permissions())

	// 注册
	router.POST("info", controller.UserInfo)

}
