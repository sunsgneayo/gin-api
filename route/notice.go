package route

import (
	"dpj-admin-api/api/controller"
	"dpj-admin-api/api/middleware"
	"github.com/gin-gonic/gin"
)

func NoticeRouter(r *gin.Engine) *gin.Engine {

	r.POST("notice", middleware.JwtAuth(), controller.GetNoticeList)

	r.POST("noticeInfo", middleware.JwtAuth(), middleware.Casbin(), controller.GetNotice)

	return r
}
