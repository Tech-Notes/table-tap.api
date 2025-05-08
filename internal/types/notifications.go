package types

import "gopkg.in/guregu/null.v4"

type NotificationMetaData struct {
	TableID int64      `json:"table_id"`
	OrderID null.Int `json:"order_id"`
}

type NotificationType string

const (
	NotificationTypeNewOrder          NotificationType = "new_order"
	NotificationTypeUpdateOrderStatus NotificationType = "update_order_status"
)

type Notification struct {
	Message    string               `json:"message" db:"message"`
	Type       NotificationType     `json:"type" db:"type"`
	IsRead     bool                 `json:"is_read" db:"is_read"`
	MetaData   NotificationMetaData `json:"meta_data" db:"meta_data"`
	BusinessID int64                `json:"business_id" db:"business_id"`
	CreatedAt  string               `json:"created_at" db:"created_at"`
	UpdatedAt  string               `json:"updated_at" db:"updated_at"`
}
