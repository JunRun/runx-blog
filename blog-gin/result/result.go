/**
 *
 * @Description: 统一返回结果
 * @Version: 1.0.0
 * @Date: 2020-01-09 10:07
 */
package result

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct {
	Code    int         `json:"code"`
	Message interface{} `json:"meg"`
	Data    interface{} `json:"data"`
}

func Success(ctx *gin.Context, data interface{}) {
	r := &Response{
		Code:    0,
		Message: "",
		Data:    data,
	}
	ctx.JSON(http.StatusOK, r)
}

func Error(ctx *gin.Context, msg string) {
	r := &Response{
		Code:    1,
		Message: msg,
		Data:    nil,
	}
	ctx.JSON(http.StatusBadRequest, r)
}

func AuthError(ctx *gin.Context, msg string) {
	r := &Response{
		Code:    2,
		Message: msg,
		Data:    nil,
	}
	ctx.JSON(http.StatusForbidden, r)
}
