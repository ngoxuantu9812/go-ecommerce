package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-ecomm/internal/service"
	"go-ecomm/internal/vo"
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
	var params vo.UserRegistrationRequest
	if err := c.ShouldBind(&params); err != nil {
		response.ErrorResponse(c, "Params are invalid")
	}
	fmt.Printf("Email params: %s", params)
	result := uc.userService.Register(params.Email, params.Purpose)
	response.HTTPStatusCodeSuccess(c, result, nil)
}
