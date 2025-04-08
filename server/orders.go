package main

import (
	"database/sql"
	"errors"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	internal "github.com/table-tap/api/internal/types"
	types "github.com/table-tap/api/internal/types"
)

type GetOrdersResponse struct {
	Orders []*types.Order `json:"orders"`
}

type GetOrdersSuccessResponse struct {
	*internal.ResponseBase
	Data *GetOrdersResponse `json:"data"`
}

func GetOrdersByTableIDHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	businessID := BusinessIDFromContext(ctx)

	tableIDSring := chi.URLParam(r, "table_id")
	tableID, err := strconv.ParseInt(tableIDSring, 10, 64)

	if err != nil {
		writeError(w, http.StatusBadRequest, errors.New("table_id is required"))
		return
	}

	orders, err := DBConn.GetOrdersByTableID(ctx, businessID, tableID)
	if err != nil && err != sql.ErrNoRows {
		writeError(w, http.StatusInternalServerError, errors.New("failed to get orders"))
		return
	}

	writeJSON(w, http.StatusOK, GetOrdersSuccessResponse{
		ResponseBase: &internal.SuccessResponse,
		Data: &GetOrdersResponse{
			Orders: orders,
		},
	})
}

func GetBusinessOrdersHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	businessID := BusinessIDFromContext(ctx)

	orders, err := DBConn.GetBusinessOrders(ctx, businessID)
	if err != nil && err != sql.ErrNoRows {
		writeError(w, http.StatusInternalServerError, errors.New("failed to get orders"))
		return
	}

	writeJSON(w, http.StatusOK, GetOrdersSuccessResponse{
		ResponseBase: &internal.SuccessResponse,
		Data: &GetOrdersResponse{
			Orders: orders,
		},
	})
}