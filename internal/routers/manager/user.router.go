package manager

import (
	"github.com/gin-gonic/gin"
)

type UserRouter struct{}

func (ur UserRouter) InitUserRouter(Router *gin.RouterGroup) {
	userRouterPublic := Router.Group("/admin/user")
	{
		userRouterPublic.POST("/active_user")
	}

}
