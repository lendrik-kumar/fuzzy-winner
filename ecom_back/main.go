package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/lendrik-kumar/ecom-back/controllers"
	"github.com/lendrik-kumar/ecom-back/database"
	"github.com/lendrik-kumar/ecom-back/middlewares"
	"github.com/lendrik-kumar/ecom-back/routes"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	app := controllers.NewApplication(database.ProductData(database.Client, "Products"), database.userData(database.Client, "Users"))

	router := gin.New()
	router.Use(gin.Logger())

	routes.UserRoutes(routes)
	routes.Use(middlewares.Authentication())

	router.GET("/addtocart", app.AddToCart())
	router.GET("/removeitem", app.RemoveItem())
	router.GET("/cartcheckout", app.BuyFromCart())
	router.GET("/instantbuy", app.InstantBuy())

	log.Fatal(router.Run(":", port))
}
