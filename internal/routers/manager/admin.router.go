package manager

import (
	"github.com/gin-gonic/gin"
)

type AdminRouter struct{}

func (ar AdminRouter) InitAdminRouter(Router *gin.RouterGroup) {
	userRouterPublic := Router.Group("/admin/")
	{
		userRouterPublic.POST("/login")
	}

}
