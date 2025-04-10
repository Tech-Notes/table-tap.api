package main

import (
	"database/sql"
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
	internal.ResponseBase
	Data *GetOrdersResponse `json:"data"`
}

func GetOrdersByTableIDHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	businessID := BusinessIDFromContext(ctx)

	tableIDSring := chi.URLParam(r, "table_id")
	tableID, err := strconv.ParseInt(tableIDSring, 10, 64)

	if err != nil {
		writeError(w, http.StatusBadRequest, ErrRequiredTableID)
		return
	}

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

func GetBusinessOrdersHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	businessID := BusinessIDFromContext(ctx)

	orders, err := DBConn.GetBusinessOrders(ctx, businessID)
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

type GetOrderDetailByIDResponse struct {
	Order *types.OrderDetail `json:"order"`
}
type GetOrderDetailByIDSuccessResponse struct {
	internal.ResponseBase
	Data *GetOrderDetailByIDResponse `json:"data"`
}

func GetOrderDetailByIDHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	businessID := BusinessIDFromContext(ctx)

	orderIDSring := chi.URLParam(r, "order_id")
	orderID, err := strconv.ParseInt(orderIDSring, 10, 64)

	if err != nil {
		writeError(w, http.StatusBadRequest, ErrRequiredOrderID)
		return
	}

	orderDetail, err := DBConn.GetOrderDetailByID(ctx, businessID, orderID)
	if err != nil && err != sql.ErrNoRows {
		writeError(w, http.StatusInternalServerError, err)
		return
	}

	writeJSON(w, http.StatusOK, GetOrderDetailByIDSuccessResponse{
		ResponseBase: internal.SuccessResponse,
		Data: &GetOrderDetailByIDResponse{
			Order: orderDetail,
		},
	})
}

type ChangeOrderStatusRequest struct {
	Status types.OrderStatus `json:"status"`
}

type ChangeOrderStatusResponse struct {
	ID int64 `json:"id"`
}

type ChangeOrderStatusSuccessResponse struct {
	internal.ResponseBase
	Data *ChangeOrderStatusResponse `json:"data"`
}

func ChangeOrderStatusHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	businessID := BusinessIDFromContext(ctx)

	orderIDSring := chi.URLParam(r, "order_id")
	orderID, err := strconv.ParseInt(orderIDSring, 10, 64)

	if err != nil {
		writeError(w, http.StatusBadRequest, ErrRequiredOrderID)
		return
	}

	data := &ChangeOrderStatusRequest{}
	err = readJSON(r, data)
	if err != nil {
		writeError(w, http.StatusInternalServerError, ErrFailedRequestBody)
		return
	}

	if data.Status == "" || !data.Status.IsValid() {
		writeError(w, http.StatusBadRequest, ErrInvalidOrderStatus)
		return
	}

	err = DBConn.ChangeOrderStatus(ctx, businessID, orderID, data.Status)
	if err != nil {
		writeError(w, http.StatusInternalServerError, err)
		return
	}

	writeJSON(w, http.StatusOK, ChangeOrderStatusSuccessResponse{
		ResponseBase: internal.SuccessResponse,
		Data: &ChangeOrderStatusResponse{
			ID: orderID,
		},
	})
}
