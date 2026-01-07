package main

import (
	"github.com/gin-gonic/gin"
	"github.com/lendrik-kumar/otp_worker/api"
)

func main() {
	router := gin.Default()

	app := api.Config{Router : router}
		
	app.Routes()
	
	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "OK",
		})
	})
	
	router.Run(":8000")
	
}