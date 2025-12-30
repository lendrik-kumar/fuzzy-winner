package websocket

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Pool struct {
	Register   chan *Client
	Unregister chan *Client
	Clients    map[*Client]bool
	Broadcast  chan Message
	Collection *mongo.Collection 
}

func NewPool(col *mongo.Collection) *Pool {
	return &Pool{
		Register:   make(chan *Client),
		Unregister: make(chan *Client),
		Clients:    make(map[*Client]bool),
		Broadcast:  make(chan Message),
		Collection: col,
	}
}

func (p *Pool) Start() {
	for {
		select {
		case client := <-p.Register:
			p.Clients[client] = true
			fmt.Println("Size of Connection Pool: ", len(p.Clients))
			// for c := range p.Clients {
			//	  c.Conn.WriteJSON(Message{Type: 1, Body: "New User Joined..."})
			// }
			
		case client := <-p.Unregister:
			delete(p.Clients, client)
			fmt.Println("Size of Connection Pool: ", len(p.Clients))
			// for c := range p.Clients {
			// 	c.Conn.WriteJSON(Message{Type: 1, Body: "User Disconnected..."})
			// }
		
		case message := <-p.Broadcast:
			
			go func(msg Message) {
				ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
				defer cancel()
			
					_, err := p.Collection.InsertOne(ctx, bson.M{
						"body":      msg.Body,
						"type":      msg.Type,
						"timestamp": time.Now(),
					})
					if err != nil {
						log.Printf("Mongo save error: %v", err)
					}
				
			}(message)
			
			for client := range p.Clients {
				if err := client.Conn.WriteJSON(message); err != nil {
					fmt.Printf("Broadcast error to client: %v\n", err)
					client.Conn.Close()
					delete(p.Clients, client)
				}
			}
		}
	}
}