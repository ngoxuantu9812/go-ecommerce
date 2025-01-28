package repo

import (
	"go-ecomm/global"
	"go-ecomm/internal/model"
)

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
	row := global.Mdb.Table(TableNameGoCrmUser).Where("usr_email = ?", email).First(&model.GoCrmUser{}).RowsAffected
	return row != NumberNull
}

func NewUserRepository() IUserRepository {
	return &userRepository{}
}
