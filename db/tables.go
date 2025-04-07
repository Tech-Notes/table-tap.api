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