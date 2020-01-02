/**
 *
 * @Description:
 * @Version: 1.0.0
 * @Date: 2019-12-20 14:51
 */
package controller

import (
	"github.com/JunRun/blog-gin/model"
	"github.com/gin-gonic/gin"
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
		Name:     c.Param("name"),
		Password: c.Param("password"),
		Address:  "",
		Auth:     "admin",
	}

}
