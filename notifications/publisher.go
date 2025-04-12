package notifications

import (
    "context"
    "fmt"
)

func (s *Server) PublishOrder(businessID, message string) error {
    channel := fmt.Sprintf("new:orders.business:%s", businessID)
    return s.Redis.Publish(context.Background(), channel, message).Err()
}