package shopper

import (
	"github.com/go-chi/chi"
	"github.com/table-tap/api/notifications"
)

func GetRouter() *chi.Mux {
	shopper := chi.NewRouter()

	shopper.Use(verify)

	shopper.Route("/menu_items", func(menuItems chi.Router) {
		menuItems.Get("/", GetMenuItemsHandler)
	})

	shopper.Route("/orders", func(r chi.Router) {
		r.Post("/", CreateOrderHandler)
	})

	//notifications
	shopper.Route("/notifications", func(noti chi.Router) {
		noti.Get("/", notifications.WebSocketHandler(NotificationHub))
	})

	r := chi.NewRouter()
	r.Mount("/api/v1", shopper)

	return r
}
