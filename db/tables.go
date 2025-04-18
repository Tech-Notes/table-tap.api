package db

import (
	"context"

	"github.com/table-tap/api/internal/types"
)

func (db *DB) CreateTable(ctx context.Context, table *types.Table) (int64, error) {

	query := `
		INSERT INTO tables (business_id, qr_code_url, status, token, table_no, description)
		VALUES (:business_id, :qr_code_url, :status, :token, :table_no, :description)
		RETURNING id`

	id, err := db.InsertxContext(ctx, query, table)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (db *DB) GetTableList(ctx context.Context, businessID int64) ([]*types.Table, error) {

	query := `
		SELECT id, business_id, qr_code_url, status, token, COALESCE(table_no, 0) AS table_no, description
		FROM tables
		WHERE business_id = $1`

	var tables []*types.Table
	err := db.SelectContext(ctx, &tables, query, businessID)
	if err != nil {
		return nil, err
	}

	return tables, nil
}

func (db *DB) GetTableByToken(ctx context.Context, token string) (*types.Table, error) {
	query := `
	SELECT id,
	business_id,
	status,
	qr_code_url,
	COALESCE(table_no, 0) AS table_no,
	description
	FROM tables
	WHERE token = $1
	`
	table := &types.Table{}
	err := db.GetContext(ctx, table, query, token)
	if err != nil {
		return nil, err
	}

	return table, nil
}

func (db *DB) GetTableByID(ctx context.Context, businesID, id int64) (*types.Table, error) {
	query := `
	SELECT id,
	business_id,
	status,
	qr_code_url,
	token,
	COALESCE(table_no, 0) AS table_no,
	description
	FROM tables
	WHERE id = $1 AND business_id = $2
	`
	table := &types.Table{}
	err := db.GetContext(ctx, table, query, id, businesID)
	if err != nil {
		return nil, err
	}

	return table, nil
}

func (db *DB) MarkTableOrdersAsPaid(ctx context.Context, businessID, tableID int64) (int64, error) {
	tx := db.MustBeginContext(ctx)
	query := `
	UPDATE tables
	SET status = :status
	WHERE id = :id AND business_id = :business_id;
	`
	args := map[string]any{
		"status":      types.TableStatusAvailable,
		"id":          tableID,
		"business_id": businessID,
	}

	_, err := tx.NamedExecContext(ctx, query, args)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	updateOrdersStatus := `
	UPDATE orders
	SET status = :status
	WHERE table_id = :table_id
	AND status != 'paid' AND business_id = :business_id
	`
	orderArgs := map[string]any{
		"status":      types.OrderStatusPaid,
		"table_id":    tableID,
		"business_id": businessID,
	}
	_, err = tx.NamedExecContext(ctx, updateOrdersStatus, orderArgs)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	err = tx.Commit()

	return tableID, err
}
