package controller

import (
	"github.com/gin-gonic/gin"
	"go-ecomm/internal/service"
	"go-ecomm/response"
)

type UserController struct {
	userService *service.UserService
}

func NewUserController() *UserController {
	return &UserController{
		userService: service.NewUserService(),
	}
}

func (uc UserController) GetAllUser(c *gin.Context) {
	response.HTTPStatusCodeSuccess(c,
		20001,
		[]string{"tu ngo", "ngo tu", "xuan tu"},
	)
}
