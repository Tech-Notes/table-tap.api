package types

import "gopkg.in/guregu/null.v4"

type MenuItem struct {
	ID          int64       `json:"id" db:"id"`
	Name        string      `json:"name" db:"name"`
	Description null.String `json:"description" db:"description"`
	Price       float32     `json:"price" db:"price"`
	BusinessID  int64       `json:"business_id" db:"business_id"`
	PhotoID     null.Int    `json:"photo_id" db:"photo_id"`
	Category    null.String `json:"category" db:"category"`
	CategoryID  null.Int    `json:"category_id" db:"category_id"`
	CreatedAt   string      `json:"created_at" db:"created_at"`
	UpdatedAt   string      `json:"updated_at" db:"updated_at"`
}

type GetMenuItemsResponse struct {
	MenuItems []*MenuItem `json:"menu_items"`
}

type GetMenuItemsSuccessResponse struct {
	ResponseBase
	Data *GetMenuItemsResponse `json:"data"`
}