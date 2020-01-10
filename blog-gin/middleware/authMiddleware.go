/**
 *
 * @Description: 权限管理中间件
 * @Version: 1.0.0
 * @Date: 2020-01-04 16:21
 */
package middleware

import (
	"fmt"
	"github.com/JunRun/blog-gin/result"
	"github.com/JunRun/blog-gin/utils"
	"github.com/casbin/casbin"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware(e *casbin.Enforcer) gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("token")
		if len(token) == 0 {
			c.Abort()
			result.AuthError(c, "token 为空")
		}
		//解析 token 判断是否存在
		_, err := utils.ParsingToken(token)
		if err != nil {
			c.Abort()
			result.AuthError(c, "验证token 失效"+err.Error())
		} else {
			fmt.Println("token 正常")
		}
		c.Next()
		//获取请求的uri
		//uri := c.Request.URL.RequestURI()
		//act := c.Request.Method
		//sub := "admin"
		//if e.Enforce(sub, uri, act) {
		//	fmt.Println("拥有权限，通过")
		//	c.Next()
		//} else {
		//	fmt.Println("没有权限，拒绝")
		//	c.Abort()
		//}

	}
}
