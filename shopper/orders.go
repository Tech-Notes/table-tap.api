package shopper

import (
	"database/sql"
	"net/http"

	internal "github.com/table-tap/api/internal/types"
)

func GetOrdersByTableIDHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	businessID := businessIDFromContext(ctx)

	tableID := tableIDFromContext(ctx)

	orders, err := DBConn.GetOrdersByTableID(ctx, businessID, tableID)
	if err != nil && err != sql.ErrNoRows {
		writeError(w, http.StatusInternalServerError, err)
		return
	}

	writeJSON(w, http.StatusOK, internal.GetOrdersSuccessResponse{
		ResponseBase: internal.SuccessResponse,
		Data: &internal.GetOrdersResponse{
			Orders: orders,
		},
	})
}