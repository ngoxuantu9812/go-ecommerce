package initialize

import (
	"go-ecomm/global"
	"go-ecomm/internal/database"
	"go-ecomm/service"
	"go-ecomm/service/impl"
)

func InitServiceInterface() {
	queries := database.New(global.Mdbc)
	// User service interface
	service.InitUserLogin(impl.NewUserLoginImpl(queries))
}
