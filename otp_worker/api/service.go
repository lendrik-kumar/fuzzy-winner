package api

import (
	"errors"

	"github.com/twilio/twilio-go"
	twilioApi "github.com/twilio/twilio-go/rest/verify/v2"
)

var client *twilio.RestClient = twilio.NewRestClientWithParams(twilio.ClientParams{
	Username : envAccountSID(),
	Password : envAuthToken(),
})

func (app *Config) twilioSendOTP(phone_number string) (string, error){
	params := &twilioApi.CreateVerificationParams{}
	params.SetTo(phone_number)
	params.SetChannel("sms")

	res, err := client.VerifyV2.CreateVerification(envTwilioSID(), params)
	if err != nil {
		return "", err
	}

	return *res.Sid, nil
}

func (app *Config) twilioVerifyOTP(phone_number, code string) error{
	params := &twilioApi.CreateVerificationCheckParams{}
	params.SetTo(phone_number)
	params.SetCode(code)

	res, err := client.VerifyV2.CreateVerificationCheck(envTwilioSID(), params)
	if err != nil {
		return err
	}
	if *res.Status != "approved" {
		return  errors.New("not valid otp")
	}

	return nil
}