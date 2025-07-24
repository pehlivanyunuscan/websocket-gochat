package websocketgochat

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	h := hub.NewHub()
	go h.Run() // Start the hub to handle client connections

	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		serveWs(h, w, r) // Handle WebSocket connections
	})
	srv := &http.Server{
		Addr: ":8080", // Set the server address
	}
	go func() {
		log.Println("Starting WebSocket server on :8080")
		if err := srv.ListenAndServe(); err != nil {
			log.Fatal("ListenAndServe:", err) // Log any errors starting the server
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM) // Listen for interrupt or terminate signals
	<-quit                                               // Wait for a signal to shut down
	log.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel() // Ensure the context is cancelled after use
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err) // Log any errors during shutdown
	}
	log.Println("Server gracefully stopped")
}
