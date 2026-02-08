//go:build wireinject

package wire

import (
	"github.com/google/wire"
	"github.com/huyle/go-ecommerce-backend-api/internal/controller"
	"github.com/huyle/go-ecommerce-backend-api/internal/repo"
	"github.com/huyle/go-ecommerce-backend-api/internal/service"
)

func InitUserRouterHandle() (*controller.UserController, error) {
	wire.Build(
		repo.NewUserRepository,
		service.NewUserService,
		controller.NewUserController,
	)
	return new(controller.UserController), nil
}
