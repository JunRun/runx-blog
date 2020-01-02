/**
 *
 * @Description:
 * @Version: 1.0.0
 * @Date: 2019-12-23 10:01
 */
package model

type UserModel struct {
	Name     string `json:"name"`
	Password string `json:"password"`
	Address  string `json:"address"`
	Auth     string `json:"auth"`
}

func (u *UserModel) Login() error {

}
