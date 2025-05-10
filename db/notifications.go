package db

import (
	"context"

	"github.com/table-tap/api/internal/types"
)

func (db *DB) GetNotificationList(ctx context.Context, businessID int64) ([]*types.Notification, error) {
	query := `
	SELECT n.id,
	n.message,
	n.type,
	n.is_read,
	n.meta_data,
	n.business_id
	FROM notifications n
	WHERE n.business_id = $1
	ORDER BY n.created_at DESC;
	`

	notifications := []*types.Notification{}

	err := db.SelectContext(ctx, &notifications, query, businessID)

	if err != nil {
		return nil, err
	}

	return notifications, nil
}

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

func (db *DB) SetNotificationAsReadByID(ctx context.Context, id, businesID int64) error {
	query := `
	UPDATE notifications
	SET is_read = TRUE
	WHERE id = $1 AND business_id = $2;
	`

	_, err := db.ExecContext(ctx, query, id, businesID)

	return err
}
