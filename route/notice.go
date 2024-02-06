package route

import (
	"dpj-admin-api/api/controller"
	"github.com/gin-gonic/gin"
)

func NoticeRouter(r *gin.RouterGroup) {

	//router := r.Group("notice").Use(middleware.Permissions())
	router := r.Group("notice")

	router.GET("noticeList", controller.GetNoticeList)

	router.POST("noticeInfo", controller.GetNotice)

}
