package route

import (
	"dpj-admin-api/api/controller"
	"dpj-admin-api/api/middleware"
	"github.com/gin-gonic/gin"
)

func RoomsRouter(r *gin.RouterGroup) {
	router := r.Group("room").Use(middleware.JwtAuth())

	// 房间列表
	router.POST("list", controller.RoomList)

}
