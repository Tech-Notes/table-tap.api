package shopper

import (
	"database/sql"
	"net/http"
	"time"

	"github.com/table-tap/api/internal/types"
	utils "github.com/table-tap/api/internal/utils"
)

func GetOrdersByTableIDHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	businessID := utils.BusinessIDFromContext(ctx)

	tableID := utils.TableIDFromContext(ctx)

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

func CreateOrderHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	businessID := utils.BusinessIDFromContext(ctx)

	tableID := utils.TableIDFromContext(ctx)

	id, err := DBConn.CreateOrder(ctx, businessID, tableID)
	if err != nil {
		writeError(w, http.StatusInternalServerError, err)
		return
	}

	// Publish new order notification
	NotificationServer.PublishOrderNotification(businessID, &types.NewOrderNotiPayload{
		TableNumber: tableID,
		OrderID:     id,
		Type:        types.NotificationTypeNewOrder,
		Status:      types.OrderStatusPending,
		CreatedAt:   time.Now().Format(time.RFC3339),
	})

	writeJSON(w, http.StatusCreated, types.ActionSuccessResponse{
		ResponseBase: types.SuccessResponse,
		Data: &types.ActionSuccessResponseData{
			ID: id,
		},
	})
}
