package wbsocket

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/johneliud/real-time-forum/database"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

var clients = make(map[*websocket.Conn]bool)

// WebSocket connection handler
func HandleWebSocket(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Error during connection upgrade:", err)
		return
	}
	defer conn.Close()

	// Register the new client
	clients[conn] = true
	log.Println("New client connected")

	for {
		var msg struct {
			Content string `json:"content"`
			Sender  string `json:"sender"`
		}
		if err := conn.ReadJSON(&msg); err != nil {
			log.Println("Error reading message:", err)
			break
		}

		// Store the message in the database
		if err := database.InsertMessage(msg.Content, msg.Sender); err != nil {
			log.Println("Error inserting message:", err)
		}

		// Broadcast the message to all connected clients
		for client := range clients {
			if err := client.WriteJSON(msg); err != nil {
				log.Println("Error sending message to client:", err)
				client.Close()
				delete(clients, client)
			}
		}
	}
}
