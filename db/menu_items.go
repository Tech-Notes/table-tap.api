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

func (db *DB) CreateMenuItem(ctx context.Context, businessID int64, item *internal.MenuItem) (int64, error) {
	query := `
	INSERT INTO menu_items (name, description, price, business_id, photo_id, category, category_id)
	VALUES (:name, :description, :price, :business_id, :photo_id, :category, :category_id)
	RETURNING id;
	`

	item.BusinessID = businessID

	id, err := db.InsertxContext(ctx, query, item)
	if err != nil {
		return 0, err
	}

	return id, nil
}
