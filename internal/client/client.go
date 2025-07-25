package client

import (
	"log"
	"time"
	"websocket-gochat/internal/hub"
	"websocket-gochat/internal/types"

	"github.com/gorilla/websocket"
)

// Client represents a WebSocket client.
type Client struct {
	Conn     *websocket.Conn
	Send     chan types.Message // send channel for messages
	Username string
}

// ReadMessages reads messages from the WebSocket connection.
func ReadMessages(c *types.Client, h *hub.Hub) {
	defer func() {
		h.Unregister <- c
		c.Conn.Close()
	}()

	c.Conn.SetReadLimit(512)                                 // Set a read limit for incoming messages
	c.Conn.SetReadDeadline(time.Now().Add(60 * time.Second)) // Set a read deadline
	c.Conn.SetPongHandler(func(string) error {               // Set a pong handler to reset the read deadline
		c.Conn.SetReadDeadline(time.Now().Add(60 * time.Second)) // Reset read deadline on pong
		return nil
	})
	for {
		var msg types.Message
		err := c.Conn.ReadJSON(&msg)
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("WebSocket error: %v", err)
			}
			break
		}
		msg.Username = c.Username // Set the username for the message
		h.Broadcast <- msg        // Broadcast the message to all clients
	}
}

// WriteMessages writes messages to the WebSocket connection.
func WriteMessages(c *types.Client) {
	ticker := time.NewTicker(54 * time.Second) // Send ping messages to keep the connection alive
	defer func() {
		ticker.Stop()
		c.Conn.Close()
	}()
	for {
		select {
		case msg, ok := <-c.Send:
			c.Conn.SetWriteDeadline(time.Now().Add(10 * time.Second))
			if !ok {
				c.Conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}
			if err := c.Conn.WriteJSON(msg); err != nil {
				log.Println("Error writing JSON:", err)
				return
			}

		case <-ticker.C:
			c.Conn.SetWriteDeadline(time.Now().Add(10 * time.Second))
			if err := c.Conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				return
			}
		}
	}
}
