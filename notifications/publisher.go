package notifications

import (
	"context"
	"encoding/json"
	"fmt"
)

func (s *Server) PublishOrderNotification(businessID int64, payload any) error {
    channel := fmt.Sprintf("orders.business:%d", businessID)
    msg, err := json.Marshal(payload)
    if err != nil {
        return err
    }
    return s.Redis.Publish(context.Background(), channel, msg).Err()
}