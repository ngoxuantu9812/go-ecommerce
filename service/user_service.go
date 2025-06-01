package service

import "context"

type (
	IUserLogin interface {
		Login(ctx context.Context) error
		Register(ctx context.Context) error
		VerifyOTP(ctx context.Context, otp string) error
		UpdatePasswordRegister(ctx context.Context) error
	}

	IUserInfo interface {
		GetInfoByUserId(ctx context.Context) error
		GetAllUser(ctx context.Context) error
	}

	IUserAdmin interface {
		RemoveUser(ctx context.Context) error
		FindOneUser(ctx context.Context) error
	}
)

var (
	localUserAdmin IUserAdmin
	localUserLogin IUserLogin
	localUserInfo  IUserInfo
)

func UserAdmin() IUserAdmin {
	if localUserAdmin == nil {
		panic("Implement LocalUserAdmin() not found interface")
	}
	return localUserAdmin
}

func InitUserAdmin(i IUserAdmin) {
	localUserAdmin = i
}

func UserLogin() IUserLogin {
	if localUserLogin == nil {
		panic("Implement LocalUserLogin() not found interface")
	}
	return localUserLogin
}

func InitUserLogin(i IUserLogin) {
	localUserLogin = i
}

func UserInfo() IUserInfo {
	if localUserInfo == nil {
		panic("Implement LocalUserInfo() not found interface")
	}
	return localUserInfo
}

func InitUserInfo(i IUserInfo) {
	localUserInfo = i
}
