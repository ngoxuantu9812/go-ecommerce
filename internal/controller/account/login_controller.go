package account

import (
	"github.com/gin-gonic/gin"
	"go-ecomm/response"
	"go-ecomm/service"
)

var Login = new(cUserLogin)

type cUserLogin struct {
}

func (c *cUserLogin) Login(ctx *gin.Context) {
	err := service.UserLogin().Login(ctx)
	if err != nil {
		response.ErrorResponse(ctx, err.Error())
		return
	}

	response.HTTPStatusCodeSuccess(ctx, response.ErrCodeSuccess, nil)

}
