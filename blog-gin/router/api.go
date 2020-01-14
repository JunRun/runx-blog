/**
 *
 * @Description:
 * @Version: 1.0.0
 * @Date: 2019-12-20 13:49
 */
package router

import (
	"github.com/JunRun/blog-gin/controller"
	"github.com/JunRun/blog-gin/middleware"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	//使用允许跨域的中间件
	r.Use(middleware.CorsMiddleware())
	//登录接口
	r.POST("/login", controller.Login)
	//注册接口
	r.POST("/register", controller.UserRegister)

	//使用群组中间件，实现鉴权认证
	user := r.Group("/user", middleware.AuthMiddleware(nil))
	{
		user.GET("/hello", controller.Hello)

	}

	r.Run(":9092")
	return r

}
