package notifications

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/redis/go-redis/v9"
	utils "github.com/table-tap/api/internal/utils"
)

type Server struct {
	Redis    *redis.Client
	Upgrader websocket.Upgrader
}

type Options struct {
	RedisAddr string
}

func NewServer(options *Options) *Server {
	rdb := redis.NewClient(&redis.Options{Addr: options.RedisAddr})

	return &Server{
		Redis: rdb,
		Upgrader: websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool { return true },
		},
	}
}

// Handles WebSocket connection for a specific business
func (s *Server) HandleWS(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	businessID := utils.BusinessIDFromContext(ctx)
	if businessID == 0 {
		log.Println("Invalid business ID")
		return
	}

	tableID := utils.TableIDFromContext(ctx)

	conn, err := s.Upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("WebSocket upgrade error:", err)
		return
	}
	defer conn.Close()

	channel := fmt.Sprintf("orders.business:%d", businessID)

	if tableID != 0 {
		channel = fmt.Sprintf("orders.table:%d", tableID)
	}
	
	sub := s.Redis.Subscribe(ctx, channel)
	defer sub.Close()

	for {
		msg, err := sub.ReceiveMessage(ctx)
		if err != nil {
			log.Println("Redis receive error:", err)
			return
		}

		if err := conn.WriteMessage(websocket.TextMessage, []byte(msg.Payload)); err != nil {
			log.Println("WebSocket write error:", err)
			return
		}
	}
}
