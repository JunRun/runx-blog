/**
 *
 * @Description:
 * @Version: 1.0.0
 * @Date: 2020-01-04 09:58
 */
package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

//跨域中间件
func CorsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		//origin:=c.Request.Header.Get("origin")
		//var fliterHost = [...]string{"http://localhost.*"}
		var accentsAble = false
		if accentsAble {
			// 核心处理方式
			c.Header("Access-Control-Allow-Origin", "*")
			c.Header("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept")
			c.Header("Access-Control-Allow-Methods", "GET, OPTIONS, POST, PUT, DELETE")
			c.Set("content-type", "application/json")
		}
		//放行所有OPTIONS方法
		if method == "OPTIONS" {
			c.JSON(http.StatusOK, "Options Request!")
		}

		c.Next()

	}
}
