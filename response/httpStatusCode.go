package response

const (
	ErrCodeSuccess      = 20001
	ErrCodeParamInvalid = 20003
	ErrTokenInvalid     = 30001
	ErrOtpInvalid       = 30002
	ErrSendEmailOtp     = 30003
	ErrUserHasExist     = 50001
)

var msg = map[int]string{
	ErrCodeSuccess:      "Success",
	ErrCodeParamInvalid: "Email is invalid",
	ErrTokenInvalid:     "Token is invalid",
	ErrUserHasExist:     "User has exist",
	ErrOtpInvalid:       "Otp is invalid",
	ErrSendEmailOtp:     "Send email error",
}
