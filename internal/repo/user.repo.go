package repo

// type UserRepo struct{}

// func NewUserRepo() *UserRepo {
// 	return &UserRepo{}
// }

// // GetInfoUser retrieves user information from the repository
// func (ur *UserRepo) GetInfoUser() string {
// 	return "User Info from Repo"
// }

type IUserRepository interface {
	GetUserByEmail(email string) bool
}

type userRepository struct {
}

// GetUserByEmail implements [IUserRepository].
func (u *userRepository) GetUserByEmail(email string) bool {
	panic("unimplemented")
}

func NewUserRepository() IUserRepository {
	return &userRepository{}
}
