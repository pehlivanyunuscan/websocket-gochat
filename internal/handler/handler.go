package handler

import (
	"log"
	"net/http"
	"websocket-gochat/internal/client"
	"websocket-gochat/internal/hub"
	"websocket-gochat/internal/types"

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

	c := &types.Client{
		Conn:     conn,
		Send:     make(chan types.Message, 256), // Buffered channel for sending messages
		Username: "",                            // Use the initial message as the username
	}

	go client.ReadMessages(c, h) // Start reading messages from the client
	go client.WriteMessages(c)   // Start writing messages to the client
}
