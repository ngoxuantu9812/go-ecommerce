package routers

import (
	"go-ecomm/internal/routers/manager"
	"go-ecomm/internal/routers/user"
)

type RouterGroup struct {
	User    user.UserRouterGroup
	Manager manager.ManageRouterGroup
}

var RouterGroupApp = new(RouterGroup)
