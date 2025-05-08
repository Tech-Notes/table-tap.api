package db

import (
	"context"

	"github.com/table-tap/api/internal/types"
)

func (db *DB) CreateNotification(ctx context.Context, notification *types.Notification) (int64, error) {
	query := `
	INSERT INTO notifications (message, type, is_read, meta_data, business_id)
	VALUES (:message, :type, :is_read, :meta_data, :business_id)
	RETURNING id
	`

	orderID, err := db.InsertxContext(ctx, query, notification)
	if err != nil {
		return 0, err
	}

	return orderID, nil
}
