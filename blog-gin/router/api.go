/**
 *
 * @Description:
 * @Version: 1.0.0
 * @Date: 2019-12-20 13:49
 */
package router

import (
	"github.com/JunRun/blog-gin/controller"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	user := r.Group("/user")
	{
		user.GET("/hello", controller.Hello)
		//用户登录接口
		user.POST("/login", controller.Login)
		//
		user.POST("/register", controller.UserRegister)
	}

	r.Run(":9091")
	return r

}
