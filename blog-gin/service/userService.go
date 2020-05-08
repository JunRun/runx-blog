/**
 *
 * @Description:
 * @Version: 1.0.0
 * @Date: 2020-01-02 13:42
 */
package service

import (
	"github.com/JunRun/blog-gin/model"
	"github.com/JunRun/blog-gin/utils"
	"github.com/gin-gonic/gin"
)

type userService struct {
}

var UserService *userService

func (u *userService) Login(c *gin.Context, user *model.UserModel) error {

	if err := user.Login(); err != nil {
		return err
	} else {
		token := utils.GenerateToken(user.UserName, user.Password)
		user.Token = token
		return nil
	}

}

func (u *userService) Register(user *model.UserModel) error {

	utils.RedisClient.Get("id").Result()
	//密码加密
	user.Password = utils.MD5Encryption(user.Password)
	if err := user.Register(); err != nil {
		return err
	} else {
		return nil
	}
}
