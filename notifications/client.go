package notifications

import "github.com/gorilla/websocket"

type Client struct {
	conn   *websocket.Conn
	send   chan []byte
	topics map[string]bool
	hub    *Hub
}