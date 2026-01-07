package data

type OtpData struct {
	PhoneNumber string `json:"phone_number,omitempty" validate:"required"`
}

type VerifyData struct {
	User *OtpData `json:"user,omitempty" validate:"required"`
	OTP  string   `json:"otp,omitempty" validate:"required"`
}