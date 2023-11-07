package middleware

import (
	"dpj-admin-api/config"
	response "dpj-admin-api/support"
	"github.com/casbin/casbin/v2"
	gormandiser "github.com/casbin/gorm-adapter/v3"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Permissions casbin权限认证系统
func Permissions() gin.HandlerFunc {
	return func(c *gin.Context) {
		a, _ := gormandiser.NewAdapterByDB(config.Db()) // Your driver and data source.
		e, _ := casbin.NewEnforcer("./config/casbin/rbac_model.conf", a)
		//ip := c.ClientIP()
		uri := c.Request.URL.Path

		method := c.Request.Method

		userId, _ := c.Get("UserId")

		//e.AddPolicy(userId, uri, method)

		// Check the permission.
		result, err := e.Enforce(userId, uri, method)
		if err != nil {
			response.WithContext(c).Error(http.StatusForbidden, "认证失败")
			c.Abort()
			return
		}

		if result == false {
			response.WithContext(c).Error(http.StatusForbidden, "认证失败，无权限访问")
			c.Abort()
			return
		}
		c.Next()

	}
}
