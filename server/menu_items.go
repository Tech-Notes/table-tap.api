package main

import (
	"database/sql"
	"net/http"

	internal "github.com/table-tap/api/internal/types"
)

type GetMenuItemsResponse struct {
	MenuItems []*internal.MenuItem `json:"menu_items"`
}

type GetMenuItemsSuccessResponse struct {
	internal.ResponseBase
	Data *GetMenuItemsResponse `json:"data"`
}

func GetMenuItemsHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	businessID := BusinessIDFromContext(ctx)

	// Get menu items from the database
	menuItems, err := DBConn.GetMenuItems(r.Context(), businessID)
	if err != nil && err != sql.ErrNoRows {
		writeError(w, http.StatusInternalServerError, err)
	}

	writeJSON(w, http.StatusOK, GetMenuItemsSuccessResponse{
		ResponseBase: internal.SuccessResponse,
		Data: &GetMenuItemsResponse{
			MenuItems: menuItems,
		},
	})
}
