/**
 *
 * @Description: 权限管理中间件
 * @Version: 1.0.0
 * @Date: 2020-01-04 16:21
 */
package middleware

import (
	"fmt"
	"github.com/casbin/casbin"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware(e *casbin.Enforcer) gin.HandlerFunc {
	return func(c *gin.Context) {
		//获取请求的uri
		uri := c.Request.URL.RequestURI()
		act := c.Request.Method
		sub := "admin"
		if e.Enforce(sub, uri, act) {
			fmt.Println("拥有权限，通过")
			c.Next()
		} else {
			fmt.Println("没有权限，拒绝")
			c.Abort()
		}

	}
}
