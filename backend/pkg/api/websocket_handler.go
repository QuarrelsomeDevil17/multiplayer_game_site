package api

import (
	//"fmt"
	"log"
	"net/http"
	"github.com/gorilla/websocket"
)

// Define the WebSocket upgrader
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		// Allow connections from any origin. In production, restrict to specific origins.
		return true
	},
}

// WebSocketHandler upgrades the HTTP connection to a WebSocket connection
// and handles communication with the client
func WebSocketHandler(w http.ResponseWriter, r *http.Request) {
	// Upgrade the HTTP connection to a WebSocket connection
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Failed to upgrade connection:", err)
		http.Error(w, "Failed to upgrade connection", http.StatusInternalServerError)
		return
	}
	defer conn.Close()

	log.Println("Client connected:", r.RemoteAddr)

	// Handle incoming and outgoing messages
	for {
		// Read message from the WebSocket
		messageType, msg, err := conn.ReadMessage()
		if err != nil {
			log.Println("Error reading message:", err)
			break
		}

		log.Printf("Received message: %s\n", msg)

		// Here, you can implement your game logic to process the message.
		// For now, we're just echoing the message back to the client.
		err = conn.WriteMessage(messageType, msg)
		if err != nil {
			log.Println("Error sending message:", err)
			break
		}
	}

	log.Println("Client disconnected:", r.RemoteAddr)
}
