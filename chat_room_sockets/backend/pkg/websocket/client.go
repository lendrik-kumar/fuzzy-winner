package websocket

import (
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

type Client struct {
	ID   string
	Conn *websocket.Conn
	Pool *Pool
	mu   sync.Mutex
}

type ChatMessage struct {
	Username  string    `bson:"username" json:"username"`
	Body      string    `bson:"body" json:"body"`
	Timestamp time.Time `bson:"timestamp" json:"timestamp"`
}

type Message struct {
	Type int    `json:"type"`
	Body string `json:"body"`
}

func (c *Client) Read() {
	defer func() {
		c.Pool.Unregister <- c
		c.Conn.Close()
	}()

	for {
		messageType, p, err := c.Conn.ReadMessage()
		if err != nil {
			log.Println("read error:", err)
			return
		}

		message := Message{
			Type: messageType,
			Body: string(p),
		}

		c.Pool.Broadcast <- message

		fmt.Printf("message received: %+v\n", message)
	}
}
