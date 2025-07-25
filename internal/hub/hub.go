package hub

import (
	"log"
	"sync"
	"websocket-gochat/internal/types"
)

type Hub struct {
	Clients    map[*types.Client]bool
	Broadcast  chan types.Message
	Register   chan *types.Client
	Unregister chan *types.Client
	mu         sync.Mutex
}

func NewHub() *Hub {
	return &Hub{
		Clients:    make(map[*types.Client]bool),
		Broadcast:  make(chan types.Message),
		Register:   make(chan *types.Client),
		Unregister: make(chan *types.Client),
	}
}

func (h *Hub) Run() {
	for {
		select {
		case client := <-h.Register: // Register a new client
			h.mu.Lock()
			h.Clients[client] = true
			h.mu.Unlock()
			log.Printf("Client %s connected", client.Username)

		case client := <-h.Unregister: // Unregister a client
			h.mu.Lock()
			if _, ok := h.Clients[client]; ok { // Check if client is registered
				delete(h.Clients, client)
				close(client.Send)
				log.Printf("Client %s disconnected", client.Username)
			}
			h.mu.Unlock()

		case message := <-h.Broadcast: // Broadcast a message to all clients
			h.mu.Lock()
			log.Printf("Broadcasting message from %s: %s", message.Username, message.Content)
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
