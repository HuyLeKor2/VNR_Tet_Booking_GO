package manager

import "github.com/gin-gonic/gin"

type AdminRouter struct {
}

func (r *AdminRouter) InitAdminRouter(Router *gin.RouterGroup) {
	//public routes
	AdminRouterPublic := Router.Group("admin"){
		UserRouterPublic.POST("/login", nil)
	}

	//private routes
	AdminRouterPublic := Router.Group("admin/user")
	// AdminRouterPublic.Use(Limiter()
	// AdminRouterPublic.Use(Auth())
	// AdminRouterPublic.Use(Permission())
	{
		AdminRouterPublic.POST("/active_user", nil)
	}
}
