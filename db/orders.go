package db

import (
	"context"

	internal "github.com/table-tap/api/internal/types"
)

func (db *DB) GetBusinessOrders(ctx context.Context, businessID int64) ([]*internal.Order, error) {
	query := `
		SELECT o.id,
		o.business_id, 
		o.table_id, 
		o.status
		FROM orders o
		WHERE o.business_id = $1
	`
	orders := []*internal.Order{}
	err := db.SelectContext(ctx, &orders, query, businessID)
	if err != nil {
		return nil, err
	}
	return orders, nil
}

func (db *DB) GetOrdersByTableID(ctx context.Context, businessID, tableID int64) ([]*internal.Order, error) {
	query := `
		SELECT o.id,
		o.business_id, 
		o.table_id, 
		o.status
		FROM orders o
		WHERE o.table_id = $1
		AND o.business_id = $2
	`
	orders := []*internal.Order{}
	err := db.SelectContext(ctx, &orders, query, tableID, businessID)
	if err != nil {
		return nil, err
	}
	return orders, nil
}
