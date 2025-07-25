package hub

import (
	"sync"
	"websocket-gochat/message"

	"websocket-gochat/internal/client"
)

type Hub struct {
	Clients    map[*client.Client]bool
	Broadcast  chan message.Message
	Register   chan *client.Client
	Unregister chan *client.Client
	mu         sync.Mutex
}

func NewHub() *Hub {
	return &Hub{
		Clients:    make(map[*client.Client]bool),
		Broadcast:  make(chan message.Message),
		Register:   make(chan *client.Client),
		Unregister: make(chan *client.Client),
	}
}
func (h *Hub) Run() {
	for {
		select {
		case client := <-h.Register: // Register a new client
			h.mu.Lock()
			h.Clients[client] = true
			h.mu.Unlock()

		case client := <-h.Unregister: // Unregister a client
			h.mu.Lock()
			if _, ok := h.Clients[client]; ok { // Check if client is registered
				delete(h.Clients, client)
				close(client.Send)
			}
			h.mu.Unlock()

		case message := <-h.Broadcast: // Broadcast a message to all clients
			h.mu.Lock()
			for client := range h.Clients {
				select {
				case client.Send <- message:
				default:
					close(client.Send)
					delete(h.Clients, client)
				}
			}
			h.mu.Unlock()
		}
	}
}
