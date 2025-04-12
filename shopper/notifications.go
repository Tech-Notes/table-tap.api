package shopper

import "net/http"

func NewOrderNotificationHandler(w http.ResponseWriter, r *http.Request) {
	NotificationServer.HandleWS(w, r)
}