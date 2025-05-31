package service

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/segmentio/kafka-go"
	"go-ecomm/global"
	"go-ecomm/internal/repo"
	"go-ecomm/internal/untils/crypto"
	"go-ecomm/internal/untils/random"
	"go-ecomm/response"
	"strconv"
	"time"
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

type userService struct {
	userRepo     repo.IUserRepository
	userAuthRepo repo.IUserAuthRepository
}

func NewUserService(
	userRepo repo.IUserRepository,
	userAuthRepo repo.IUserAuthRepository,
) IUserService {
	return &userService{
		userRepo:     userRepo,
		userAuthRepo: userAuthRepo,
	}
}

func (us *userService) Register(email string, purpose string) int {

	//0. Hash email
	hashEmail := crypto.GetHash(email)
	fmt.Printf("hashEmail:::%s", hashEmail)
	//5. Check OTP is available

	//6. user spam ...

	//1. Check email exists in db
	if us.userRepo.GetUserByEmail(email) {
		return response.ErrUserHasExist
	}

	//2. new OTP
	otp := random.GenerateSixDigitOtp()
	if purpose == "TEST_USER" {
		otp = 123456
	}

	fmt.Printf("OTP is :::%d\n", otp)

	//3. save OTP in redis with expiration time
	err := us.userAuthRepo.AddOTP(hashEmail, strconv.Itoa(otp), int64(10*time.Minute))
	if err != nil {
		return response.ErrOtpInvalid
	}
	//4. send email otp
	//err = sendto.SendTemplateEmailOTp(
	//	[]string{email},
	//	"hello@demomailtrap.com",
	//	"otp-auth.html",
	//	map[string]interface{}{
	//		"otp": strconv.Itoa(otp),
	//	})

	// send email otp via kafka

	body := make(map[string]interface{})
	body["email"] = email
	body["otp"] = otp
	bodyRequest, _ := json.Marshal(body)
	message := kafka.Message{
		Key:   []byte("otp-auth"),
		Value: bodyRequest,
		Time:  time.Now(),
	}

	err = global.KafkaProducer.WriteMessages(context.Background(), message)

	if err != nil {
		return response.ErrSendEmailOtp
	}

	return response.ErrCodeSuccess
}
