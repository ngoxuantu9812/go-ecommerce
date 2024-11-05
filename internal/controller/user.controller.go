package controller

import (
	"github.com/gin-gonic/gin"
	"go-ecomm/internal/service"
	"go-ecomm/response"
)

type UserController struct {
	userService service.IUserService
}

func NewUserController(
	userService service.IUserService,
) *UserController {
	return &UserController{
		userService: userService,
	}
}

func (uc *UserController) GetAllUser(c *gin.Context) {
	response.HTTPStatusCodeSuccess(c,
		20001,
		[]string{"tu ngo", "ngo tu", "xuan tu"},
	)
}

func (uc *UserController) Register(c *gin.Context) {
	result := uc.userService.Register("", "")
	response.HTTPStatusCodeSuccess(c, result, nil)
}
