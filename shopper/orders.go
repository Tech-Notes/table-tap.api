package shopper

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

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

	// Create the notification payload
	notification := map[string]string{
		"code":    "new_order",
		"message": "A new order has been created",
		"orderID": fmt.Sprintf("%d", id),
	}

	// Convert the notification to a JSON string
	notificationJSON, err := json.Marshal(notification)
	if err != nil {
		writeError(w, http.StatusInternalServerError, err)
		return
	}

	// Publish new order notification to admin
	NotificationHub.Publish("admin", notificationJSON)

	writeJSON(w, http.StatusCreated, types.ActionSuccessResponse{
		ResponseBase: types.SuccessResponse,
		Data: &types.ActionSuccessResponseData{
			ID: id,
		},
	})
}
