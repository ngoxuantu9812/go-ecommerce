package impl

import (
	"context"
	"go-ecomm/internal/database"
)

type SUserLogin struct {
	r *database.Queries
}

func NewUserLoginImpl(r *database.Queries) *SUserLogin {
	return &SUserLogin{
		r: r,
	}
}

func (s *SUserLogin) Login(ctx context.Context) error {
	return nil
}

func (s *SUserLogin) Register(ctx context.Context) error {
	return nil
}

func (s *SUserLogin) VerifyOTP(ctx context.Context, otp string) error {
	return nil
}

func (s *SUserLogin) UpdatePasswordRegister(ctx context.Context) error {
	return nil
}
