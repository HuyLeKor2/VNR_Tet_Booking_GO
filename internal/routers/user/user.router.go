package user

import (
	"github.com/gin-gonic/gin"
	"github.com/huyle/go-ecommerce-backend-api/internal/controller"
	"github.com/huyle/go-ecommerce-backend-api/internal/repo"
	"github.com/huyle/go-ecommerce-backend-api/internal/service"
)

type UserRouter struct {
}

func (r *UserRouter) InitUserRouter(Router *gin.RouterGroup) {
	userController, _ := wire.InitUserRouterHandle()
	//WIRE

	UserRouterPublic := Router.Group("/user"){
		UserRouterPublic.POST("/register", userController.Register) //register -> yes/no)
	}
}

//register ->
// verify email/phone -> otp -> set password
// UserRouterPublic.POST("/otp", nil)

//private routes

// UserRouterPublic.Use(Limiter()
// UserRouterPublic.Use(AuNewUserController()iddleware())
// UserRouterPublic.Use(Permissione())
