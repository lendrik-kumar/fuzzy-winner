package api

import (
	"github.com/gin-gonic/gin"
)

type Config struct {
	Router *gin.Engine
}

func (app *Config) Routes() {
	app.Router.POST("/send-otp", app.SendOTP())
	app.Router.POST("/verify-otp", app.VerifyOTP())
}