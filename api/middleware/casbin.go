package middleware

import (
	"fmt"
	"github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"github.com/gin-gonic/gin"
)

func Casbin() gin.HandlerFunc {
	return func(c *gin.Context) {
		a, _ := gormadapter.NewAdapter("mysql", "data_center:KCMBfAjeJhbJXsSe@tcp(43.138.132.9:3398)/") // Your driver and data source.
		e, _ := casbin.NewEnforcer("./config/rbac_model.conf", a)

		e.LoadPolicy()

		e.AddPolicy("alice", "data1", "read")
		// Check the permission.
		ok, err := e.Enforce("alice", "data1", "wi")
		if err != nil {
			fmt.Printf("%s", err)
		}

		if ok == true {
			fmt.Printf("ok")
			// 允许alice读取data1
		} else {
			fmt.Printf("认证失败")

			c.JSON(403, gin.H{
				"message": "认证失败，无权限访问",
			})
			c.Abort()
			return
		}

	}
}
