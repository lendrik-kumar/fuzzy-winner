package api

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/lendrik-kumar/otp_worker/data"
)

var appTimeout = time.Second * 10

func (app *Config) SendOTP() gin.HandlerFunc {
	return func(c *gin.Context) {

		_, cancel := context.WithTimeout(context.Background(), appTimeout)
		defer cancel()

		var payload data.OtpData

		if err := app.validateBody(c, &payload); err != nil {
			app.errorJSON(c, err)
			return
		}

		newData := data.OtpData{
			PhoneNumber: payload.PhoneNumber,
		}

		_, err := app.twilioSendOTP(newData.PhoneNumber)
		if err != nil {
			app.errorJSON(c, err)
			return
		}

		app.writeJSON(c, http.StatusAccepted, gin.H{
			"message": "OTP sent successfully",
		})
	}
}

func (app *Config) VerifyOTP() gin.HandlerFunc {
	return func(c *gin.Context) {

		_, cancel := context.WithTimeout(context.Background(), appTimeout)
		defer cancel()

		var payload data.VerifyData

		if err := app.validateBody(c, &payload); err != nil {
			app.errorJSON(c, err)
			return
		}

		newData := data.VerifyData{
			User: payload.User,
			OTP:  payload.OTP,
		}

		err := app.twilioVerifyOTP(
			newData.User.PhoneNumber,
			newData.OTP,
		)

		if err != nil {
			fmt.Println("err:", err)
			app.errorJSON(c, err)
			return
		}

		app.writeJSON(c, http.StatusAccepted, gin.H{
			"message": "OTP verified successfully",
		})
	}
}
