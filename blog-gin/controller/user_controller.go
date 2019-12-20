/**
 *
 * @Description:
 * @Version: 1.0.0
 * @Date: 2019-12-20 14:51
 */
package controller

import "github.com/gin-gonic/gin"

type UserController struct {
}

func Hello(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "hello",
	})
}
