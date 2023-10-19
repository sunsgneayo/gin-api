package route

import "github.com/gin-gonic/gin"

func InitRoute(r *gin.Engine) {

	CommonRouter(r)
	NoticeRouter(r)
}
