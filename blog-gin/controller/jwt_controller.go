/**
 *
 * @Description:
 * @Version: 1.0.0
 * @Date: 2020-01-04 15:37
 */
package controller

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"time"
)

type Params struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func GetToken(c *gin.Context) {
	var params Params
	err := c.BindJSON(&params)
	if err != nil {
		fmt.Println("发生了某些错误", err)
		return
	}
	hmacSampleSecret := []byte("my_secret_ky")
	//进行 算法加密
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": params.Username,
		"password": params.Password,
		"nbf":      time.Date(2020, 01, 04, 12, 0, 0, 0, time.UTC).Unix(),
	})
	tokenString, err := token.SignedString(hmacSampleSecret)
	fmt.Println(tokenString, err)
	c.JSON(200, gin.H{
		"token": tokenString,
	})
}
