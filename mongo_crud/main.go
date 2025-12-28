package main

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"
	"github.com/lendrik-kumar/mongo-crud/controllers"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func getMongoClient() *mongo.Client {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(
		"",//mongo connection string
	))

	if err != nil {
		log.Fatal(err)
	}

	return client
}

func main() {
	r := httprouter.New()
	client := getMongoClient()

	uc := controllers.NewUserController(client)

	r.GET("/user/:id", uc.GetUser)
	r.POST("/user", uc.CreateUser)
	r.PUT("/user/:id", uc.UpdateUser)
	r.DELETE("/user/:id", uc.DeleteUser)

	log.Println("Server running on :8000")
	log.Fatal(http.ListenAndServe(":8000", r))
}