//go:build wireinject

package wire

import (
	"github.com/google/wire"
	"go-ecomm/internal/controller"
	"go-ecomm/internal/repo"
	"go-ecomm/internal/service"
)

func InitUserRouterHandle() (*controller.UserController, error) {
	wire.Build(
		service.NewUserService,
		repo.NewUserRepository,
		controller.NewUserController,
	)
	return new(controller.UserController), nil
}
