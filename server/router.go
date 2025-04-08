package main

import "github.com/go-chi/chi"

func GetRootRouter() *chi.Mux {
	r := chi.NewRouter()

	r.Mount("/api", getApiRouter())

	return r
}

func getApiRouter() *chi.Mux {
	api := chi.NewRouter()

	api.Post("/v1/signin", SignInHandler)

	v1 := chi.NewRouter()

	v1.Use(verify)

	v1.Route("/healthCheck", func(r chi.Router) {
		r.Get("/", authorizeHandler(DashboardView, HealthCheckHandler))
	})

	v1.Route("/menu_items", func(menuItems chi.Router) {
		menuItems.Get("/", authorizeHandler(DashboardView, GetMenuItemsHandler))
	})

	v1.Route("/tables", func(tables chi.Router) {
		tables.Post("/", authorizeHandler(CreateTable, CreateTableHandler))
		tables.Get("/", authorizeHandler(DashboardView, GetTablesHandler))
	})

	api.Mount("/v1", v1)

	return api
}
