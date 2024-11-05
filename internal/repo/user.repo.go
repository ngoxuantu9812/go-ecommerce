package repo

//type UserRepo struct{}
//
//func NewUserRepo() *UserRepo {
//	return &UserRepo{}
//}
//
//func (ur *UserRepo) GetInfoUser() string {
//	return "Hello World"
//}

// Interface

type IUserRepository interface {
	GetUserByEmail(email string) bool
}

type userRepository struct{}

func (ur *userRepository) GetUserByEmail(email string) bool {
	return true
}

func NewUserRepository() IUserRepository {
	return &userRepository{}
}
