package notifications

import "sync"


type Hub struct {
	clients    map[*Client]bool
	topics     map[string]map[*Client]bool
	register   chan *Client
	unregister chan *Client
	subscribe  chan subscription
	publish    chan publication
	mu         sync.Mutex
}

type subscription struct {
	client *Client
	topic  string
}

type publication struct {
	topic   string
	message []byte
}

// Hub maintains everything
func NewHub() *Hub {
	return &Hub{
		clients:    make(map[*Client]bool),
		topics:     make(map[string]map[*Client]bool),
		register:   make(chan *Client),
		unregister: make(chan *Client),
		subscribe:  make(chan subscription),
		publish:    make(chan publication),
	}
}


func (h *Hub) Run() {
	for {
		select {
		case client := <-h.register:
			h.clients[client] = true

		case client := <-h.unregister:
			delete(h.clients, client)
			for topic := range client.topics {
				delete(h.topics[topic], client)
				if len(h.topics[topic]) == 0 {
					delete(h.topics, topic)
				}
			}
			close(client.send)

		case sub := <-h.subscribe:
			if sub.client.topics == nil {
				sub.client.topics = make(map[string]bool)
			}
			sub.client.topics[sub.topic] = true

			if h.topics[sub.topic] == nil {
				h.topics[sub.topic] = make(map[*Client]bool)
			}
			h.topics[sub.topic][sub.client] = true

		case pub := <-h.publish:
			if clients, ok := h.topics[pub.topic]; ok {
				for client := range clients {
					select {
					case client.send <- pub.message:
					default:
						close(client.send)
						delete(h.clients, client)
					}
				}
			}
		}
	}
}

// --- External Methods ---

func (h *Hub) Publish(topic string, message []byte) {
	h.publish <- publication{topic: topic, message: message}
}

func (h *Hub) Subscribe(client *Client, topic string) {
	h.subscribe <- subscription{client: client, topic: topic}
}

func (h *Hub) Register(client *Client) {
	h.register <- client
}

func (h *Hub) Unregister(client *Client) {
	h.unregister <- client
}