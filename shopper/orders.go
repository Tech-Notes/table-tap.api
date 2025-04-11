package shopper

import (
	"database/sql"
	"net/http"

	internal "github.com/table-tap/api/internal/types"
)

type GetOrdersResponse struct {
	Orders []*internal.Order `json:"orders"`
}

type GetOrdersSuccessResponse struct {
	internal.ResponseBase
	Data *GetOrdersResponse `json:"data"`
}

func GetOrdersByTableIDHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	businessID := businessIDFromContext(ctx)

	tableID := tableIDFromContext(ctx)

	orders, err := DBConn.GetOrdersByTableID(ctx, businessID, tableID)
	if err != nil && err != sql.ErrNoRows {
		writeError(w, http.StatusInternalServerError, err)
		return
	}

	writeJSON(w, http.StatusOK, GetOrdersSuccessResponse{
		ResponseBase: internal.SuccessResponse,
		Data: &GetOrdersResponse{
			Orders: orders,
		},
	})
}