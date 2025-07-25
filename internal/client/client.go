package client

import (
	"log"
	"time"
	"websocket-gochat/internal/hub"
	"websocket-gochat/message"

	"github.com/gorilla/websocket"
)

type Client struct {
	Conn     *websocket.Conn
	Send     chan message.Message // send channel for messages
	Username string
}

func (c *Client) ReadMessages(h *hub.Hub) {
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
		var msg message.Message
		err := c.Conn.ReadJSON(&msg)
		if err != nil {
			log.Println("Error reading JSON:", err)
			return
		}
		msg.Username = c.Username // Set the username for the message
		h.Broadcast <- msg        // Broadcast the message to all clients
	}
}

func (c *Client) WriteMessages() {
	ticker := time.NewTicker(60 * time.Second) // Send ping messages to keep the connection alive
	defer func() {
		ticker.Stop()
		c.Conn.Close()
	}()
	for {
		select {
		case msg, ok := <-c.Send:
			if !ok {
				c.Conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}
			if err := c.Conn.WriteJSON(msg); err != nil {
				log.Println("Error writing JSON:", err)
				return
			}

		case <-ticker.C:
			if err := c.Conn.WriteMessage(websocket.PingMessage, []byte{}); err != nil {
				log.Println("Error sending ping:", err)
				return
			}
		}
	}
}
