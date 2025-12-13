package shopper

import (
	"database/sql"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/table-tap/api/internal/types"
)

func GetMenuItemsHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	shopIDString := chi.URLParam(r, "id")
	shopID, err := strconv.ParseInt(shopIDString, 10, 64)
	if err != nil {
		writeError(w, http.StatusBadRequest, ErrRequiredShopID)
		return
	}

	// Get menu items from the database
	menuItems, err := DBConn.GetMenuItems(ctx, shopID)
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
