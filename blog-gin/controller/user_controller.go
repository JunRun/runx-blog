/**
 *
 * @Description:
 * @Version: 1.0.0
 * @Date: 2019-12-20 14:51
 */
package controller

import (
	"github.com/JunRun/blog-gin/model"
	"github.com/JunRun/blog-gin/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserController struct {
}

func Hello(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "hello",
	})
}

func Login(c *gin.Context) {
	user := &model.UserModel{
		Name:     c.GetString("username"),
		Password: c.Query("password"),
		Address:  "",
		Auth:     "admin",
	}
	if err := service.UserService.Login(user); err != nil {
		c.JSON(http.StatusFound, err.Error())
	} else {
		c.JSON(http.StatusOK, "登录成功")
	}
}

func UserRegister(c *gin.Context) {
	user := &model.UserModel{
		Id:       0,
		Name:     c.Query("username"),
		Password: c.Query("password"),
		Address:  "",
		Auth:     "",
	}
	if err := service.UserService.Register(user); err != nil {
		c.JSON(http.StatusForbidden, err.Error())
	} else {
		c.JSON(http.StatusOK, "注册成功")
	}
}
