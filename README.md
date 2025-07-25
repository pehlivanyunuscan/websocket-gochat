# WebSocket GoChat 🚀

A real-time chat application built with Go and WebSockets using Gorilla WebSocket library. This project demonstrates concurrent programming, WebSocket communication, and clean architecture principles in Go.

## ✨ Features

- **Real-time messaging** - Instant message delivery using WebSockets
- **Multiple clients support** - Handle multiple concurrent connections
- **Username authentication** - Initial message requires username
- **System notifications** - Join/leave chat notifications
- **Graceful shutdown** - Proper cleanup on server termination
- **Concurrent architecture** - Goroutines for scalable performance
- **Thread-safe operations** - Mutex protection for shared resources
- **Connection management** - Ping/pong keepalive mechanism
- **Error handling** - Comprehensive error management
- **Modular design** - Clean separation of concerns


### Core Components

- **Hub**: Central message broadcaster and client manager
- **Client**: Individual WebSocket connection handler
- **Handler**: HTTP/WebSocket upgrade handler
- **Types**: Shared data structures (Message, Client)

## 📋 Prerequisites

- Go 1.24.4 or higher
- Git (for cloning)
- WebSocket client for testing (wscat, browser, etc.)

## 🚀 Installation

### 1. Clone the repository
```bash
git clone <repository-url>
cd websocket-gochat
```

### 2. Install dependencies
```bash
go mod download
go mod tidy
```

### 3. Run the application
```bash
go run main.go
```
