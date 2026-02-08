package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/huyle/go-ecommerce-backend-api/internal/service"
	"github.com/huyle/go-ecommerce-backend-api/pkg/response"
)

type UserController struct {
	UserService service.IUserService
}

func NewUserController(us service.IUserService) *UserController {
	return &UserController{
		UserService: us,
	}
}

func (uc *UserController) Register(gin *gin.Context) int {
	result := uc.UserService.Register("", "")
	response.SuccessResponse(gin, result, nil)
	return result
}

// // controller -> service -> repo -> database
// func (uc *UserController) GetUserByID(c *gin.Context) {
// 	response.SuccessResponse(c, 20001, []string{"M10", "huylekore"})
// }
