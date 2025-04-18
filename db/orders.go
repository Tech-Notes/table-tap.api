package db

import (
	"context"
	"database/sql"

	"github.com/table-tap/api/internal/types"
)

func (db *DB) GetBusinessOrders(ctx context.Context, businessID int64) ([]*types.Order, error) {
	query := `
		SELECT o.id,
		o.business_id, 
		o.table_id, 
		COALESCE(o.table_no, 0) AS table_no,
		o.status
		FROM orders o
		WHERE o.business_id = $1
		ORDER BY id DESC;
	`
	orders := []*types.Order{}
	err := db.SelectContext(ctx, &orders, query, businessID)
	if err != nil {
		return nil, err
	}
	return orders, nil
}

func (db *DB) GetOrdersByTableID(ctx context.Context, businessID, tableID int64) ([]*types.Order, error) {
	query := `
		SELECT o.id,
		o.business_id, 
		o.table_id,
		COALESCE(o.table_no, 0) AS table_no, 
		o.status
		FROM orders o
		WHERE o.table_id = $1
		AND o.business_id = $2
		AND o.status != 'paid'
		ORDER BY id DESC;
	`
	orders := []*types.Order{}
	err := db.SelectContext(ctx, &orders, query, tableID, businessID)
	if err != nil {
		return nil, err
	}
	return orders, nil
}

func (db *DB) GetOrderDetailByID(ctx context.Context, businessID, orderID int64) (*types.OrderDetail, error) {
	query := `
		SELECT o.id,
		o.business_id, 
		o.table_id, 
		COALESCE(o.table_no, 0) AS table_no,
		o.status
		FROM orders o
		WHERE o.business_id = $1
		AND o.id = $2
	`
	orderDetail := &types.OrderDetail{}
	err := db.GetContext(ctx, orderDetail, query, businessID, orderID)
	if err != nil {
		return nil, err
	}

	orderItemsQuery := `
		SELECT oi.id,
		oi.item_id,
		oi.order_id,
		oi.quantity,
		oi.price
		FROM orders_items oi
		WHERE oi.order_id = $1
	`
	orderItems := []*types.OrderItem{}
	err = db.SelectContext(ctx, &orderItems, orderItemsQuery, orderID)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}
	orderDetail.OrderItems = orderItems

	return orderDetail, nil
}

func (db *DB) ChangeOrderStatus(ctx context.Context, businessID, orderID int64, status types.OrderStatus) error {
	query := `
	UPDATE orders
	SET status = $1
	WHERE id = $2 AND business_id = $3
	`

	_, err := db.ExecContext(ctx, query, status, orderID, businessID)
	if err != nil {
		return err
	}

	return nil
}

func (db *DB) CreateOrder(ctx context.Context, businessID, tableID int64) (int64, error) {
	query := `
	INSERT INTO orders (business_id, table_id, status)
	VALUES (:business_id, :table_id, :status)
	RETURNING id
	`

	args := map[string]interface{}{
		"business_id": businessID,
		"table_id":    tableID,
		"status":      types.OrderStatusPending,
	}

	orderID, err := db.InsertxContext(ctx, query, args)
	if err != nil {
		return 0, err
	}

	return orderID, nil
}

func (db *DB) GetOrderByID(ctx context.Context, businessID, orderID int64) (*types.Order, error) {
	query := `
		SELECT o.id,
		o.business_id, 
		o.table_id, 
		COALESCE(o.table_no, 0) AS table_no,
		o.status
		FROM orders o
		WHERE o.business_id = $1
		AND o.id = $2
	`
	order := &types.Order{}
	err := db.GetContext(ctx, order, query, businessID, orderID)
	if err != nil {
		return nil, err
	}
	return order, nil
}
