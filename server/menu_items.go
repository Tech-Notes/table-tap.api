package main

import (
	"database/sql"
	"net/http"

	internal "github.com/table-tap/api/internal/types"
)

func GetMenuItemsHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	businessID := BusinessIDFromContext(ctx)

	// Get menu items from the database
	menuItems, err := DBConn.GetMenuItems(r.Context(), businessID)
	if err != nil && err != sql.ErrNoRows {
		writeError(w, http.StatusInternalServerError, err)
	}

	writeJSON(w, http.StatusOK, internal.GetMenuItemsSuccessResponse{
		ResponseBase: internal.SuccessResponse,
		Data: &internal.GetMenuItemsResponse{
			MenuItems: menuItems,
		},
	})
}

type CreateMenuItemResponse struct {
	ID int64 `json:"id"`
}

type CreateMenuItemSuccessResponse struct {
	internal.ResponseBase
	Data *CreateMenuItemResponse `json:"data"`
}

func CreateMenuItemHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	businessID := BusinessIDFromContext(ctx)

	item := &internal.MenuItem{}
	err := readJSON(r, item)
	if err != nil {
		writeError(w, http.StatusBadRequest, ErrFailedRequestBody)
		return
	}

	id, err := DBConn.CreateMenuItem(ctx, businessID, item)
	if err != nil {
		writeError(w, http.StatusInternalServerError, err)
		return
	}

	writeJSON(w, http.StatusCreated, CreateMenuItemSuccessResponse{
		ResponseBase: internal.SuccessResponse,
		Data: &CreateMenuItemResponse{
			ID: id,
		},
	})
}
