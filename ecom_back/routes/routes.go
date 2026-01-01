package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/lendrik-kumar/ecom-back/controllers"
)

func userRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.POST("users/signup", controllers.signup())
	incomingRoutes.POST("users/login", controllers.login())
	incomingRoutes.POST("/admin/addproduct", controllers.addproduct())
	incomingRoutes.GET("/users/search", controllers.searchProductByquery())
	incomingRoutes.GET("/users/productview", controllers.searchProduct())
}
