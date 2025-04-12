package main

import (
	"database/sql"
	"net/http"

	"github.com/table-tap/api/internal/types"
	utils "github.com/table-tap/api/internal/utils"
)

func GetMenuItemsHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	businessID := utils.BusinessIDFromContext(ctx)

	// Get menu items from the database
	menuItems, err := DBConn.GetMenuItems(r.Context(), businessID)
	if err != nil && err != sql.ErrNoRows {
		writeError(w, http.StatusInternalServerError, err)
	}

	writeJSON(w, http.StatusOK, types.GetMenuItemsSuccessResponse{
		ResponseBase: types.SuccessResponse,
		Data: &types.GetMenuItemsResponse{
			MenuItems: menuItems,
		},
	})
}

func CreateMenuItemHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	businessID := utils.BusinessIDFromContext(ctx)

	item := &types.MenuItem{}
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

	writeJSON(w, http.StatusCreated, types.ActionSuccessResponse{
		ResponseBase: types.SuccessResponse,
		Data: &types.ActionSuccessResponseData{
			ID: id,
		},
	})
}
