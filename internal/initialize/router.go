package initialize

import (
	"github.com/gin-gonic/gin"
	"go-ecomm/global"
	"go-ecomm/internal/routers"
	"go-ecomm/response"
)

func InitRouter() *gin.Engine {
	var r *gin.Engine
	if global.Config.Sever.Mode == "dev" {
		gin.SetMode(gin.DebugMode)
		gin.ForceConsoleColor()
		r = gin.Default()
	} else {
		gin.SetMode(gin.ReleaseMode)
		r = gin.New()
	}

	r.Use(gin.Recovery()) // logging
	r.Use(gin.Recovery()) // cross
	r.Use(gin.Recovery()) // limiter global
	manageRouter := routers.RouterGroupApp.Manager
	userRouter := routers.RouterGroupApp.User

	MainGroup := r.Group("/v1/2024")
	{
		MainGroup.GET("/check-status", func(c *gin.Context) {
			response.HTTPStatusCodeSuccess(c,
				20001,
				[]string{"tu ngo", "ngo tu", "xuan tu"},
			)
		})
	}
	{
		userRouter.InitUserRouter(MainGroup)
		userRouter.InitProductRouter(MainGroup)
	}
	{
		manageRouter.InitAdminRouter(MainGroup)
		manageRouter.InitUserRouter(MainGroup)
	}

	return r
}
