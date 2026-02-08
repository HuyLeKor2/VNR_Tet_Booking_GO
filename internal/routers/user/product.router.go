package user

import "github.com/gin-gonic/gin"

type ProductRouter struct {
}

func (r *ProductRouter) InitProductRouter(Router *gin.RouterGroup) {
	//public routes
	ProductRouterPublic := Router.Group("product"){
		ProductRouterPublic.GET("search", nil)
		ProductRouterPublic.GET("detail/:id", nil)
	}

	//private routes
}
