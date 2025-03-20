package wbsocket

import (
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/johneliud/real-time-forum/backend/logger"
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
		logger.Error("Error during connection upgrade:", "err", err)
		return
	}
	defer conn.Close()

	// Register the new client
	clients[conn] = true

	for {
		var msg struct {
			Content string `json:"content"`
			Sender  string `json:"sender"`
		}
		if err := conn.ReadJSON(&msg); err != nil {
			logger.Error("Error reading message:", "err", err)
			break
		}

		// Store the message in the database
		if err := database.InsertMessage(msg.Content, msg.Sender); err != nil {
			logger.Error("Error inserting message:", "err", err)
		}

		// Broadcast the message to all connected clients
		for client := range clients {
			if err := client.WriteJSON(msg); err != nil {
				logger.Error("Error sending message to client:", "err", err)
				client.Close()
				delete(clients, client)
			}
		}
	}
}
