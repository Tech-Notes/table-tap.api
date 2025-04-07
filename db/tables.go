package db

import (
	"context"

	internal "github.com/table-tap/api/internal/types"
)

func (db *DB) CreateTable(ctx context.Context, table *internal.Table) (int64, error) {
	
	query := `
		INSERT INTO tables (business_id, qr_code_url, status, token)
		VALUES (:business_id, :qr_code_url, :status, :token)
		RETURNING id`

	id, err := db.InsertxContext(ctx, query, table)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (db *DB) GetTableList(ctx context.Context, businessID int64) ([]*internal.Table, error) {
	
	query := `
		SELECT id, business_id, qr_code_url, status, token
		FROM tables
		WHERE business_id = $1`

	var tables []*internal.Table
	err := db.SelectContext(ctx, &tables, query, businessID)
	if err != nil {
		return nil, err
	}

	return tables, nil
}