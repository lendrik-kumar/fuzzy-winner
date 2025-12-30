package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/lendrik-kumar/chat-room/pkg/websocket"
	"github.com/rs/cors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func serveWs(pool *websocket.Pool, w http.ResponseWriter, r *http.Request) {
	conn, err := websocket.Upgrade(w, r)
	if err != nil {
		fmt.Println("upgrade err : ", err)
		return
	}

	client := &websocket.Client{
		Conn: conn,
		Pool: pool,
	}

	pool.Register <- client
	go client.Read()
}

func corsMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		next(w, r)
	}
}

func setupRoutes(r *mux.Router, pool *websocket.Pool) {
	r.HandleFunc("/ws", corsMiddleware(func(w http.ResponseWriter, r *http.Request) {
		serveWs(pool, w, r)
	}))
}

func getMongoClient() *mongo.Client {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(
		os.Getenv("MONGO_URL"),
	))

	if err != nil {
		log.Fatal(err)
	}

	return client
}

func main() {
	fmt.Println("Chat server starting on :8000...")

	err := godotenv.Load()
    if err != nil {
        log.Fatal("Error loading .env file")
    }

	mongoURL := os.Getenv("MONGO_URL")
	if mongoURL == "" {
		log.Fatal("MONGO_URL not set in environment")
	}
	
	r := mux.NewRouter()

	client := getMongoClient()
	collection := client.Database("chat_app").Collection("messages")
	
	pool := websocket.NewPool(collection)
    go pool.Start()

	setupRoutes(r, pool)

	r.HandleFunc("/messages", func(w http.ResponseWriter, r *http.Request) {

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		if err := client.Ping(ctx, nil); err != nil {
			log.Fatal("MongoDB ping failed:", err)
		}

		cursor, err := collection.Find(ctx, bson.M{})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		var messages []websocket.ChatMessage
		if err := cursor.All(ctx, &messages); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(messages)
	}).Methods("GET")

	handler := cors.Default().Handler(r)

	http.ListenAndServe(":8000", handler)
}