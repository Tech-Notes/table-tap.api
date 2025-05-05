package shopper

import (
	"github.com/table-tap/api/db"
	"github.com/table-tap/api/notifications"
)

var (
	DBConn *db.DB
	NotificationServer *notifications.Server
	NotificationHub *notifications.Hub
)