package types

import "github.com/gorilla/websocket"

type Message struct {
	Username string `json:"username"`
	Content  string `json:"content"`
}

type Client struct {
	Conn     *websocket.Conn
	Send     chan Message // send channel for messages
	Username string
}
