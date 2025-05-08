package main

import (
	"database/sql"
	"net/http"

	"github.com/table-tap/api/internal/types"
	"github.com/table-tap/api/internal/utils"
)

type NotificationListResponse struct {
	Notifications []*types.Notification `json:"notifications"`
}

type NotificationListSuccessResponse struct {
	types.ResponseBase
	Data *NotificationListResponse `json:"data"`
}

func GetNotificationListHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	businessID := utils.BusinessIDFromContext(ctx)

	notificaitons, err := DBConn.GetNotificationList(ctx, businessID)
	if err != nil && err != sql.ErrNoRows {
		writeError(w, http.StatusInternalServerError, err)
		return
	}

	writeJSON(w, http.StatusOK, NotificationListSuccessResponse{
		ResponseBase: types.SuccessResponse,
		Data: &NotificationListResponse{
			Notifications: notificaitons,
		},
	})
}
