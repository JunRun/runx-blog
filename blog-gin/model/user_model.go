/**
 *
 * @Description:
 * @Version: 1.0.0
 * @Date: 2019-12-23 10:01
 */
package model

import (
	"errors"
	"fmt"
)

type UserModel struct {
	Id       int    `json:"id" gorm:"AUTO_INCREMENT"`
	UserName string `json:"username" gorm:"size:255;index"`
	Password string `json:"password" gorm:"size:36"`
	Address  string `json:"address"`
	Auth     string `json:"auth"`
}

func (u *UserModel) Login() error {
	g := GetConnect()
	g.Where("name = ? AND password = ?", u.UserName, u.Password).First(&u)
	fmt.Println("登录检测--用户", u.UserName)
	if u.Id != 0 {

		return nil
	} else {
		return errors.New("此账号或密码不存在")
	}

}

func (u *UserModel) Register() error {
	db := GetConnect()
	//判断用户名是否重复
	user := UserModel{}
	db.Where("name = ?", u.UserName).First(&user)
	if user.Id != 0 {
		return errors.New("用户名重复")
	}
	db.Create(&u)
	if u.Id != 0 {
		return nil
	} else {
		return errors.New("注册用户失败")
	}
}
