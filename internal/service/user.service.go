package service

import (
	"go-ecomm/internal/repo"
	"go-ecomm/response"
)

//type UserService struct {
//	userRepo *repo.UserRepo
//}
//
//func NewUserService() *UserService {
//	return &UserService{
//		userRepo: repo.NewUserRepo(),
//	}
//}
//
//func (us *UserService) GetInfoUserService() string {
//	return us.userRepo.GetInfoUser()
//}

type IUserService interface {
	Register(email string, purpose string) int
}

type UserService struct {
	userRepo repo.IUserRepository
}

func (us UserService) Register(email string, purpose string) int {
	if us.userRepo.GetUserByEmail(email) {
		return response.ErrUserHasExist
	}

	return response.ErrCodeSuccess
}

func NewUserService(userRepo repo.IUserRepository) IUserService {
	return &UserService{
		userRepo: userRepo,
	}
}
