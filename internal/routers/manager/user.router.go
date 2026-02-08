package manager

import "github.com/gin-gonic/gin"

type UserRouter struct {
}

func (r *UserRouter) InitUserRouter(Router *gin.RouterGroup) {
	//public routes
	// UserRouterPublic := Router.Group("admin/user"){
	// 	UserRouterPublic.POST("/register", nil)
	// 	UserRouterPublic.POST("/otp", nil)
	// }

	//private routes
	UserRouterPublic := Router.Group("admin/user")
	// UserRouterPublic.Use(Limiter()
	// UserRouterPublic.Use(Auth())
	// UserRouterPublic.Use(Permission())
	{
		UserRouterPublic.POST("/active_user", nil)
	}
}
