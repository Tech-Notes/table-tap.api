package shopper

import "github.com/go-chi/chi"

func GetRouter() *chi.Mux {
	shopper := chi.NewRouter()

	shopper.Use(verify)

	shopper.Route("/menu_items", func(menuItems chi.Router) {
		menuItems.Get("/", GetMenuItemsHandler)
	})

	shopper.Route("/orders", func(r chi.Router) {
		r.Post("/", CreateOrderHandler)
	})

	r := chi.NewRouter()
	r.Mount("/api/v1", shopper)

	return r
}
