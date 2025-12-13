package shopper

import (
	"github.com/go-chi/chi"
	"github.com/table-tap/api/notifications"
)

func GetRouter() *chi.Mux {
	shopper := chi.NewRouter()

	shopper.Use(verify)

	shopper.Route("/shops", func(shops chi.Router) {
		shops.Get("/{id}/menu_items", GetMenuItemsHandler)
	})

	shopper.Route("/orders", func(r chi.Router) {
		r.Post("/", CreateOrderHandler)
	})

	//notifications
	shopper.Route("/notifications", func(noti chi.Router) {
		noti.Get("/", notifications.WebSocketHandler(NotificationHub))
	})

	r := chi.NewRouter()

	r.Post("/api/v1/tables/validate", ValidateTableHandler)

	r.Mount("/api/v1", shopper)

	return r
}
