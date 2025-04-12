package db

import (
	"context"

	"github.com/table-tap/api/internal/types"
)

func (db *DB) CreateTable(ctx context.Context, table *types.Table) (int64, error) {
	
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

func (db *DB) GetTableList(ctx context.Context, businessID int64) ([]*types.Table, error) {
	
	query := `
		SELECT id, business_id, qr_code_url, status, token
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
	qr_code_url
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