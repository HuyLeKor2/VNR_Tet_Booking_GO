package routers

import (
	"github.com/huyle/go-ecommerce-backend-api/internal/routers/manager"
	"github.com/huyle/go-ecommerce-backend-api/internal/routers/user"
)

type RouterGroup struct {
	User    user.UserRouterGroup
	Manager manager.AdminRouterGroup
}

var RouterGroupApp = new(RouterGroup)
