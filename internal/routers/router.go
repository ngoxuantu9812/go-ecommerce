package routers

import (
	"github.com/gin-gonic/gin"
	"go-ecomm/internal/middlewares"
)

func InitRouter() *gin.Engine {
	r := gin.Default()
	r.Use(middlewares.AuthMiddleware())

	return r
}
