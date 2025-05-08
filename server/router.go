package main

import (
	"github.com/go-chi/chi"
	"github.com/table-tap/api/notifications"
	"github.com/table-tap/api/shopper"
)

func GetRootRouter() *chi.Mux {
	r := chi.NewRouter()

	// client's routes
	r.Mount("/s", shopper.GetRouter())

	r.Mount("/api", getApiRouter())

	return r
}

func getApiRouter() *chi.Mux {
	api := chi.NewRouter()

	api.Post("/v1/signin", SignInHandler)

	v1 := chi.NewRouter()
	v1.Use(verify)

	v1.Route("/healthCheck", func(healthCheck chi.Router) {
		healthCheck.Get("/", HealthCheckHandler)
	})

	v1.Route("/menu_items", func(menuItems chi.Router) {
		menuItems.Get("/", authorizeHandler(DashboardView, GetMenuItemsHandler))
		menuItems.Post("/", authorizeHandler(DashboardView, CreateMenuItemHandler))
	})

	v1.Route("/orders", func(orders chi.Router) {
		orders.Get("/{table_id}", authorizeHandler(DashboardView, GetOrdersByTableIDHandler))
		orders.Get("/", authorizeHandler(DashboardView, GetBusinessOrdersHandler))
		orders.Get("/{order_id}", authorizeHandler(DashboardView, GetOrderDetailByIDHandler))
		orders.Patch("/{order_id}/status", authorizeHandler(DashboardView, ChangeOrderStatusHandler))
	})

	v1.Route("/tables", func(tables chi.Router) {
		tables.Post("/", authorizeHandler(CreateTable, CreateTableHandler))
		tables.Get("/", authorizeHandler(DashboardView, GetTablesHandler))
		tables.Get("/{id}", authorizeHandler(DashboardView, GetTableByIDHandler))
		tables.Post("/{id}/paid", authorizeHandler(DashboardView, MarkTableOrdersAsPaidHandler))
	})

	//notifications
	v1.Route("/notifications", func(noti chi.Router) {
		noti.Get("/ws", notifications.WebSocketHandler(NotificationHub))
		noti.Get("/", authorizeHandler(DashboardView, GetNotificationListHandler))
	})

	api.Mount("/v1", v1)

	return api
}
