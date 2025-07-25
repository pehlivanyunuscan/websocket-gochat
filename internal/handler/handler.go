package handler

import (
	"log"
	"net/http"
	"websocket-gochat/message"

	"websocket-gochat/internal/client"
	"websocket-gochat/internal/hub"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{ // Upgrade HTTP connection to WebSocket
	CheckOrigin: func(r *http.Request) bool {
		return true // Allow all origins for simplicity, adjust as needed
	},
}

func ServeWs(h *hub.Hub, w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil) // Upgrade the HTTP connection to a WebSocket connection
	if err != nil {
		log.Println("Error upgrading connection:", err)
		return
	}
	_, msg, err := conn.ReadMessage() // Read the initial message from the client
	if err != nil {
		log.Println("Error reading initial message:", err)
		conn.Close()
		return
	}
	username := string(msg) // Use the initial message as the username
	c := &client.Client{
		Conn:     conn,
		Send:     make(chan message.Message, 256), // Buffered channel for sending messages
		Username: username,
	}
	h.Register <- c // Register the new client in the hub

	go c.ReadMessages(h) // Start reading messages from the client
	go c.WriteMessages() // Start writing messages to the client
}
