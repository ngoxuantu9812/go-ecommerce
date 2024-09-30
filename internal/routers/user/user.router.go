package user

import (
	"github.com/gin-gonic/gin"
	"go-ecomm/internal/wire"
)

type UserRouter struct{}

func (ur UserRouter) InitUserRouter(Router *gin.RouterGroup) {

	//urp := repo.NewUserRepository()
	//us := service.NewUserService(urp)
	//userHanderNonDepen := controller.NewUserController(us)

	userController, _ := wire.InitUserRouterHandle()
	//Wire goes
	// Dependency Injection (DI) DI javascript
	userRouterPublic := Router.Group("/user")
	{
		userRouterPublic.POST("/register", userController.Register)
		userRouterPublic.POST("/otp")
	}

	userRouterPrivate := Router.Group("/user")
	{
		userRouterPrivate.GET("/info")
		userRouterPrivate.POST("/otp")
	}

}
