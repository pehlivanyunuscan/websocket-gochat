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

## 🏗️ Architecture

```
┌─────────────────┐    ┌─────────────────┐    ┌─────────────────┐
│   WebSocket     │    │       Hub       │    │   WebSocket     │
│   Client A      │◄──►│   (Message      │◄──►│   Client B      │
│                 │    │   Broadcaster)  │    │                 │
└─────────────────┘    └─────────────────┘    └─────────────────┘
         │                       │                       │
         │              ┌─────────────────┐              │
         └─────────────►│   HTTP Server   │◄─────────────┘
                        │   (:8080)       │
                        └─────────────────┘
```

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
git clone <https://github.com/pehlivanyunuscan/websocket-gochat>
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

The server will start on `http://localhost:8080`

## 🎯 Usage

### Starting the Server

```bash
go run main.go
# Output: Starting WebSocket server on :8080
```

### Connecting with wscat

1. **Install wscat** (if not already installed):
```bash
npm install -g wscat
```

2. **Connect to the WebSocket server**:
```bash
wscat -c ws://localhost:8080/ws
```

3. **Send your first message** (must include username):
```json
{"username": "john", "content": "Hello everyone!"}
```

4. **Send subsequent messages** (username optional):
```json
{"content": "How is everyone doing?"}
```

### Multiple Clients

Open multiple terminals and connect different users:

**Terminal 1:**
```bash
wscat -c ws://localhost:8080/ws
{"username": "alice", "content": "Hi there!"}
```

**Terminal 2:**
```bash
wscat -c ws://localhost:8080/ws
{"username": "bob", "content": "Hello Alice!"}
```

## 📡 API Documentation

### WebSocket Endpoint

**URL**: `ws://localhost:8080/ws`

**Protocol**: WebSocket

### Message Format

#### Initial Message (Required)
```json
{
  "username": "string (required)",
  "content": "string (optional)"
}
```

#### Subsequent Messages
```json
{
  "content": "string (required)"
}
```

#### System Messages (Automatic)
```json
{
  "username": "System",
  "content": "username joined the chat"
}
```

## 📁 Project Structure

```
websocket-gochat/
├── main.go                 # Application entry point
├── go.mod                  # Go module definition
├── go.sum                  # Dependency checksums
├── internal/               # Internal packages
│   ├── types/
│   │   └── types.go       # Data structures (Message, Client)
│   ├── hub/
│   │   └── hub.go         # Message hub and client manager
│   ├── client/
│   │   └── client.go      # Client connection handlers
│   └── handler/
│       └── handler.go     # HTTP/WebSocket handlers
└── README.md              # This file
```

## 📊 Example Usage

### Server Logs
```
Starting WebSocket server on :8080
Client alice connected
Client alice registered to hub
Broadcasting message from alice : content: "Hello everyone!"
Broadcasting message from System : content: "alice joined the chat"
Client bob connected
Client bob registered to hub
Broadcasting message from System : content: "bob joined the chat"
Broadcasting message from bob : content: "Hi Alice!"
```

### WebSocket Client Session
```bash
$ wscat -c ws://localhost:8080/ws
Connected (press CTRL+C to quit)
> {"username": "alice", "content": "Hello everyone!"}
< {"username":"alice","content":"Hello everyone!"}
< {"username":"System","content":"alice joined the chat"}
< {"username":"bob","content":"Hi Alice!"}
> {"content": "How are you, Bob?"}
< {"username":"alice","content":"How are you, Bob?"}
```

## 📈 Performance

- **Concurrent Connections**: Supports multiple simultaneous connections
- **Memory Efficiency**: Buffered channels with appropriate sizes
- **CPU Efficiency**: Goroutines for parallel processing
- **Network Efficiency**: WebSocket protocol for low-latency communication

## 🚦 Graceful Shutdown

The server supports graceful shutdown via:
- `SIGINT` (Ctrl+C)
- `SIGTERM` (kill command)

Shutdown process:
1. Stop accepting new connections
2. Wait for existing connections to complete (5-second timeout)
3. Clean up resources
