package shopper

import (
	"database/sql"
	"net/http"

	internal "github.com/table-tap/api/internal/types"
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

	writeJSON(w, http.StatusOK, internal.GetMenuItemsSuccessResponse{
		ResponseBase: internal.SuccessResponse,
		Data: &internal.GetMenuItemsResponse{
			MenuItems: menuItems,
		},
	})
}
