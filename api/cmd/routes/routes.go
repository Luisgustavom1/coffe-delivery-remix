package routes

import (
	"coffe-delivery-remix/api/pkg/coffes/handlers"

	"github.com/go-chi/chi/v5"
)

func Routes() *chi.Mux {
	r := chi.NewRouter()

	r.Post("/coffes", coffes.Create)
	r.Put("/coffes/{id}", coffes.Update)
	r.Delete("/coffes/{id}", coffes.Delete)
	r.Get("/coffes", coffes.List)
	r.Get("/coffes/{id}", coffes.Get)
	r.Get("/cart", coffes.List)

	return r
}
