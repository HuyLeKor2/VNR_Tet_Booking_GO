package service

import (
	"github.com/huyle/go-ecommerce-backend-api/internal/repo"
	"github.com/huyle/go-ecommerce-backend-api/pkg/response"
)

// type UserService struct {
// 	UserRepo *repo.UserRepo
// }

// func NewUserService() *UserService {
// 	return &UserService{
// 		UserRepo: repo.NewUserRepo(),
// 	}
// }

// // GetInfoUser retrieves user information from the repository
// func (us *UserService) GetInfoUser() string {
// 	return us.UserRepo.GetInfoUser()
// }

// INTERFACE_VERSION
type IUserService interface { // lam viec voi controller
	Register(email string, purpose string) int
}

type UserService struct {
	userRepo repo.IUserRepository
	//... other repos
}

// Register implements [IUserService].
func (us *UserService) Register(email string, purpose string) int {
	//1. check email exist
	if us.userRepo.GetUserByEmail(email) {
		return response.ErrCodeUserHasExists
	}
	return response.ErrCodeSuccess
}

func NewUserService(userRepo repo.IUserRepository) IUserService {
	return &UserService{
		userRepo: userRepo,
	}
}
