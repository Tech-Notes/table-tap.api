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

	api.Mount("/v1", v1)

	return api
}
