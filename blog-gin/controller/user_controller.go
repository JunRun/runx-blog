/**
 *
 * @Description:
 * @Version: 1.0.0
 * @Date: 2019-12-20 14:51
 */
package controller

import (
	"github.com/JunRun/blog-gin/model"
	"github.com/JunRun/blog-gin/result"
	"github.com/JunRun/blog-gin/service"
	"github.com/gin-gonic/gin"
)

type UserController struct {
}

func Hello(c *gin.Context) {
	result.Success(c, "hello world")
}

func Login(c *gin.Context) {
	var user = &model.UserModel{}

	if err := c.BindJSON(user); err != nil {
		result.Error(c, err.Error())
		return
	}

	if err := service.UserService.Login(c, user); err != nil {
		result.Error(c, err.Error())
	} else {
		result.Success(c, user)
	}
}

func UserRegister(c *gin.Context) {
	user := &model.UserModel{
		Id:       0,
		UserName: c.Query("username"),
		Password: c.Query("password"),
	}
	if err := service.UserService.Register(user); err != nil {
		result.Error(c, err.Error())
	} else {
		result.Success(c, "注册成功")
	}
}
