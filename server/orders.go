package main

import (
	"database/sql"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/table-tap/api/internal/types"
	utils "github.com/table-tap/api/internal/utils"
)

func GetOrdersByTableIDHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	businessID := utils.BusinessIDFromContext(ctx)

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

	writeJSON(w, http.StatusOK, types.GetOrdersSuccessResponse{
		ResponseBase: types.SuccessResponse,
		Data: &types.GetOrdersResponse{
			Orders: orders,
		},
	})
}

func GetBusinessOrdersHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	businessID := utils.BusinessIDFromContext(ctx)

	orders, err := DBConn.GetBusinessOrders(ctx, businessID)
	if err != nil && err != sql.ErrNoRows {
		writeError(w, http.StatusInternalServerError, err)
		return
	}

	writeJSON(w, http.StatusOK, types.GetOrdersSuccessResponse{
		ResponseBase: types.SuccessResponse,
		Data: &types.GetOrdersResponse{
			Orders: orders,
		},
	})
}

type GetOrderDetailByIDResponse struct {
	Order *types.OrderDetail `json:"order"`
}
type GetOrderDetailByIDSuccessResponse struct {
	types.ResponseBase
	Data *GetOrderDetailByIDResponse `json:"data"`
}

func GetOrderDetailByIDHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	businessID := utils.BusinessIDFromContext(ctx)

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
		ResponseBase: types.SuccessResponse,
		Data: &GetOrderDetailByIDResponse{
			Order: orderDetail,
		},
	})
}

type ChangeOrderStatusRequest struct {
	Status types.OrderStatus `json:"status"`
}

func ChangeOrderStatusHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	businessID := utils.BusinessIDFromContext(ctx)

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

	if data.Status == "" || !data.Status.IsValid() || data.Status == types.OrderStatusPaid {
		writeError(w, http.StatusBadRequest, ErrInvalidOrderStatus)
		return
	}

	_, err = DBConn.GetOrderByID(ctx, businessID, orderID)
	if err != nil {
		writeError(w, http.StatusInternalServerError, err)
		return
	}

	err = DBConn.ChangeOrderStatus(ctx, businessID, orderID, data.Status)
	if err != nil {
		writeError(w, http.StatusInternalServerError, err)
		return
	}

	writeJSON(w, http.StatusOK, types.ActionSuccessResponse{
		ResponseBase: types.SuccessResponse,
		Data: &types.ActionSuccessResponseData{
			ID: orderID,
		},
	})
}
