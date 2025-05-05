package notifications

import (
	"encoding/json"
	"log"

	"github.com/gorilla/websocket"
)

type Client struct {
	conn   *websocket.Conn
	send   chan []byte
	topics map[string]bool
	hub    *Hub
}

type Message struct {
	Type    string `json:"type"`
	Topic   string `json:"topic"`
	Content string `json:"content"`
}

func NewClient(conn *websocket.Conn, hub *Hub) *Client {
	return &Client{
		conn: conn,
		send: make(chan []byte, 256),
		hub:  hub,
	}
}

func (c *Client) ReadMessages() {
	defer func() {
		c.hub.Unregister(c)
		c.conn.Close()
	}()
	for {
		_, msg, err := c.conn.ReadMessage()
		if err != nil {
			break
		}
		var m Message
		if err := json.Unmarshal(msg, &m); err != nil {
			log.Println("Invalid message:", err)
			continue
		}
		switch m.Type {
		case "subscribe":
			c.hub.Subscribe(c, m.Topic)
		case "message":
			c.hub.Publish(m.Topic, []byte(m.Content))
		default:
			log.Println("Unknown message type:", m.Type)
		}
	}
}

func (c *Client) WriteMessages() {
	defer c.conn.Close()
	for msg := range c.send {
		if err := c.conn.WriteMessage(websocket.TextMessage, msg); err != nil {
			break
		}
	}
}
