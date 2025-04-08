package db

import (
	"context"

	internal "github.com/table-tap/api/internal/types"
)

func (db *DB) GetMenuItems(ctx context.Context, businessID int64) ([]*internal.MenuItem, error) {
	query := `SELECT id,
	name, 
	description, 
	price, 
	business_id, 
	photo_id, 
	category, 
	category_id 
	FROM menu_items WHERE business_id = $1`

	menuItems := []*internal.MenuItem{}
	err := db.SelectContext(ctx, &menuItems, query, businessID)
	if err != nil {
		return nil, err
	}
	
	return menuItems, nil
}