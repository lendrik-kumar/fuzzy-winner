package controllers

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/lendrik-kumar/ecom-back/models"
	"github.com/lendrik-kumar/ecom-back/tokens"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/v2/bson"
)

func HashPassword(password string) string {

}

func verifyPassword(userPassword string, hashedPassword string) (bool, string) {

}

func signup() gin.HandlerFunc {

	return func(ctx *gin.Context) {
		var c, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		var User models.User
		if err := ctx.BindJSON(&User); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return 
		}
		
		validate := validator.New()
		valErr := validate.Struct(User)
		if valErr != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": valErr})
			log.Fatal(valErr)
		}

		count, err := UserCollection.CountDocuments(c, bson.M{"email":User.Email})
		if err != nil {
			log.Panic(err)
			ctx.JSON(http.StatusInternalServerError, gin.H{"error":err})
			return 
		}

		if count > 0 {
			ctx.JSON(http.StatusBadRequest, gin.H{"error":"user exists already"})
		}

		count, err = UserCollection.CountDocuments(c, bson.M{"phone":User.Phone})
		defer cancel()
		if err != nil {
			log.Panic(err)
			ctx.JSON(http.StatusInternalServerError, gin.H{"error":err})
			return 
		}
		if count > 0 {
			ctx.JSON(http.StatusBadRequest, gin.H{"error":"this phone number is already in use"})
			return 
		}
		
		password := HashPassword(*User.Password)
		User.Password = &password

		User.CreatedAt = time.Now()
		User.UpdatedAt = time.Now()
		User.ID = primitive.NewObjectID()
		User.UserId = User.ID.Hex()

		token, refreshToken, _ := tokens.GenerateAllTokens(*User.Email, *User.FirstName, *User.LastName, User.UserId)
		User.Token = &token
		User.RefreshToken = &refreshToken

		User.UserCart = make([]models.ProductUser, 0)
		User.AddressDetails = make([]models.Address, 0)
		User.OrderStatus = make([]models.Order, 0)
		
		_, insertErr := UserCollection.InsertOne(c, User)
		if insertErr != nil {
			msg := "user item was not created"
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": msg})
			return 
		}
		defer cancel()
		ctx.JSON(http.StatusOK, "success")
	}
}

func login() gin.HandlerFunc {

	return func(ctx *gin.Context) {
		var c, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		var User models.User
		if err := ctx.BindJSON(&User); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		var foundUser models.User
		err := UserCollection.FindOne(c, bson.M{"email": User.Email}).Decode(&foundUser)
		defer cancel()
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "email or password is incorrect"})
			return 
		}

		passwordIsValid, msg := verifyPassword(*User.Password, *foundUser.Password)
		defer cancel()
		if !passwordIsValid {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": msg})
			return
		}
		
		token, refreshToken, _ := tokens.GenerateAllTokens(*foundUser.Email, *foundUser.FirstName, *foundUser.LastName, foundUser.UserId)
		
		err = UserCollection.FindOneAndUpdate(c, bson.M{"user_id": foundUser.UserId}, bson.M{"$set":bson.M{"token":token, "refresh_token":refreshToken}}).Err()
		defer cancel()
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "login failed"})
			return
		}
		ctx.JSON(http.StatusOK, foundUser)
	}
}

func productViewerAdmin() gin.HandlerFunc {

}

func searchProducts() gin.HandlerFunc {

}

func searchProductByquery() gin.HandlerFunc {

}
