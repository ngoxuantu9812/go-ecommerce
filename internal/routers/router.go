package routers

import (
	"github.com/gin-gonic/gin"
	"go-ecomm/internal/controller"
	"go-ecomm/internal/middlewares"
)

func InitRouter() *gin.Engine {
	r := gin.Default()
	r.Use(middlewares.AuthMiddleware())
	v1 := r.Group("/api/v1")
	{
		v1.GET("/ping", controller.NewUserController().GetAllUser)

	}

	return r
}
