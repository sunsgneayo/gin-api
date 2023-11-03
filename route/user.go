package route

import (
	"dpj-admin-api/api/controller"
	"github.com/gin-gonic/gin"
)

func UsersRouter(r *gin.RouterGroup) {
	router := r.Group("user")
	//router := r.Group("user").Use(middleware.JwtAuth())

	// 用户列表
	router.POST("list", controller.UserList)

}
