// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package wire

import (
	"go-ecomm/internal/controller"
	"go-ecomm/internal/repo"
	"go-ecomm/internal/service"
)

// Injectors from user.wire.go:

func InitUserRouterHandle() (*controller.UserController, error) {
	iUserRepository := repo.NewUserRepository()
	iUserService := service.NewUserService(iUserRepository)
	userController := &controller.UserController{
		userService: iUserService,
	}
	return userController, nil
}
