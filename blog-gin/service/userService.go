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
	"net/http"
	"time"
)

type userService struct {
}

var UserService *userService

func (u *userService) Login(c *gin.Context, user *model.UserModel) error {

	if err := user.Login(); err != nil {
		return err
	} else {
		cookie := &http.Cookie{
			Name:       "session_id",
			Value:      utils.GenerateToken(user.UserName, user.Password),
			Path:       "/",
			Domain:     "",
			Expires:    time.Time{},
			RawExpires: "",
			MaxAge:     0,
			Secure:     false,
			HttpOnly:   true,
			SameSite:   0,
			Raw:        "",
			Unparsed:   nil,
		}
		http.SetCookie(c.Writer, cookie)
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
