package routers

import (
	"github.com/gin-gonic/gin"
	"go-ecomm/internal/controller"
)

func NewRouter() *gin.Engine {
	r := gin.Default()
	v1 := r.Group("/api/v1")
	{
		v1.GET("/ping", controller.NewUserController().GetAllUser)

	}

	return r
}
