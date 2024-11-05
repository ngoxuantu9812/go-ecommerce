package user

import (
	"github.com/gin-gonic/gin"
	"go-ecomm/internal/wire"
)

type UserRouter struct{}

func (ur UserRouter) InitUserRouter(Router *gin.RouterGroup) {

	userController, _ := wire.InitUserRouterHandle()

	//userController, _ := wire.InitUserRouterHandle()
	//Wire goes
	// Dependency Injection (DI) DI javascript
	userRouterPublic := Router.Group("/user")
	{
		userRouterPublic.POST("/register", userController.Register)
	}

	userRouterPrivate := Router.Group("/user")
	{
		userRouterPrivate.GET("/info")
	}

}
