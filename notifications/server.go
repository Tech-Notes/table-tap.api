package notifications

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/redis/go-redis/v9"
	types "github.com/table-tap/api/internal/types"
)

type Server struct {
    Redis    *redis.Client
    Upgrader websocket.Upgrader
}

func NewServer(redisAddr string) *Server {
    rdb := redis.NewClient(&redis.Options{Addr: redisAddr})

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
	businessID := ctx.Value(types.ContextKeyBusinessID)
	businessID, ok := businessID.(int64)

    if !ok || businessID == 0 {
		log.Println("Invalid business ID")
        return
    }

    conn, err := s.Upgrader.Upgrade(w, r, nil)
    if err != nil {
        log.Println("WebSocket upgrade error:", err)
        return
    }
    defer conn.Close()

    channel := fmt.Sprintf("new:orders.business:%s", businessID)
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
