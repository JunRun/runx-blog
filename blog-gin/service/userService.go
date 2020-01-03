/**
 *
 * @Description:
 * @Version: 1.0.0
 * @Date: 2020-01-02 13:42
 */
package service

import "github.com/JunRun/blog-gin/model"

type userService struct {
}

var UserService *userService

func (u *userService) Login(user *model.UserModel) error {
	//用户密码加密处理
	user.Password = user.Password + "sda"
	err := user.Login()

	if err != nil {
		return err
	} else {
		return nil
	}
}

func (u *userService) Register(user *model.UserModel) error {
	user.Password = user.Password + "sda"
	if err := user.Register(); err != nil {
		return err
	} else {
		return nil
	}
}
