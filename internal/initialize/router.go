package initialize

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-ecomm/internal/controller"
	"go-ecomm/internal/middlewares"
)

func AA() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("Before ---> AAA")
		c.Next()
		fmt.Println("After ---> AAA")
	}
}

func BB() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("Before ---> BB")
		c.Next()
		fmt.Println("After ---> BB")
	}
}

func CC(c *gin.Context) {
	fmt.Println("Before ---> CC")
	c.Next()
	fmt.Println("After ---> C")
}

func InitRouter() *gin.Engine {
	r := gin.Default()
	r.Use(middlewares.AuthMiddleware(), AA(), BB(), CC)
	v1 := r.Group("/api/v1")
	{
		v1.GET("/ping", controller.NewUserController().GetAllUser)

	}

	return r
}
