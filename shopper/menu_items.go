package shopper

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
	menuItems, err := DBConn.GetMenuItems(ctx, businessID)
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
