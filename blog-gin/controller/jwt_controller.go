/**
 *
 * @Description:
 * @Version: 1.0.0
 * @Date: 2020-01-04 15:37
 */
package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
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

}
