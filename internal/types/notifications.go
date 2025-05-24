package types

import (
	"database/sql/driver"
	"encoding/json"
	"errors"

	"gopkg.in/guregu/null.v4"
)

type NotificationMetaData struct {
	TableID int64    `json:"table_id"`
	OrderID null.Int `json:"order_id"`
}

func (md *NotificationMetaData) Scan(value any) error {
	if value == nil {
		return nil
	}

	b, ok := value.([]byte)
	if !ok {
		return ErrInvalidNotificationMetaData
	}

	return json.Unmarshal(b, md)
}

func (md *NotificationMetaData) Value() (driver.Value, error) {
	return json.Marshal(md)
}

var ErrInvalidNotificationMetaData = errors.New("Invalid notification meta data.")

type NotificationType string

const (
	NotificationTypeNewOrder          NotificationType = "new_order"
	NotificationTypeUpdateOrderStatus NotificationType = "update_order_status"
)

type Notification struct {
	ID         int64                `json:"id" db:"id"`
	Message    string               `json:"message" db:"message"`
	Type       NotificationType     `json:"type" db:"type"`
	IsRead     bool                 `json:"is_read" db:"is_read"`
	MetaData   NotificationMetaData `json:"meta_data" db:"meta_data"`
	BusinessID int64                `json:"business_id" db:"business_id"`
	CreatedAt  string               `json:"created_at" db:"created_at"`
	UpdatedAt  string               `json:"updated_at" db:"updated_at"`
}
