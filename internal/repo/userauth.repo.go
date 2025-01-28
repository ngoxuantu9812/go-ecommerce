package repo

import (
	"fmt"
	"go-ecomm/global"
	"time"
)

type IUserAuthRepository interface {
	AddOTP(email string, otp string, expirationTime int64) error
}

type userAuthRepository struct{}

func (u *userAuthRepository) AddOTP(email string, otp string, expirationTime int64) error {
	//panic("implement me")
	key := fmt.Sprintf("usr:%s:otp", email)
	return global.Rds.Set(ctx, key, otp, time.Duration(expirationTime)).Err()
}

func NewUserAuthRepository() IUserAuthRepository {
	return &userAuthRepository{}
}
