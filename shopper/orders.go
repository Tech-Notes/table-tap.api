package shopper

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/table-tap/api/internal/types"
	utils "github.com/table-tap/api/internal/utils"
	"gopkg.in/guregu/null.v4"
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
	tableNo := utils.TableNoFromContext(ctx)

	id, err := DBConn.CreateOrder(ctx, businessID, tableID, tableNo)
	if err != nil {
		writeError(w, http.StatusInternalServerError, err)
		return
	}

	// Save notification to database
	notification := &types.Notification{
		Type:    types.NotificationTypeNewOrder,
		Message: fmt.Sprintf("A new order has been created from table - #%d", tableID),
		IsRead:  false,
		MetaData: types.NotificationMetaData{
			TableID: tableID,
			OrderID: null.IntFrom(id),
		},
		BusinessID: businessID,
	}

	_, err = DBConn.CreateNotification(ctx, notification)
	if err != nil {
		writeError(w, http.StatusInternalServerError, err)
		return
	}

	// Create the message payload
	message := map[string]string{
		"code":     "new_order",
		"message":  fmt.Sprintf("A new order has been created from table - #%d", tableID),
		"order_id": fmt.Sprintf("%d", id),
		"table_id": fmt.Sprintf("%d", tableID),
	}

	// Convert the notification to a JSON string
	notificationJSON, err := json.Marshal(message)
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
