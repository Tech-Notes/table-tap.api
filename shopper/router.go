package shopper

import "github.com/go-chi/chi"

func GetRouter() *chi.Mux {
	shopper := chi.NewRouter()

	r := chi.NewRouter()

	r.Mount("/api/v1", shopper)

	return r
}