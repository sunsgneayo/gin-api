package route

import (
	"dpj-admin-api/api/controller"
	"dpj-admin-api/api/middleware"
	"github.com/gin-gonic/gin"
)

func NoticeRouter(r *gin.RouterGroup) {

	router := r.Group("notice").Use(middleware.Permissions())

	router.POST("notice", controller.GetNoticeList)

	router.POST("noticeInfo", controller.GetNotice)

}
