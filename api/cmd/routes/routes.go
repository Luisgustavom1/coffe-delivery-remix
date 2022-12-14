package routes

import (
	cart "coffe-delivery-remix/api/pkg/controller/cart/handlers"
	coffes "coffe-delivery-remix/api/pkg/controller/coffes/handlers"

	"github.com/go-chi/chi/v5"
)

func Routes() *chi.Mux {
	r := chi.NewRouter()

	coffeRoutes(r)
	cartRoutes(r)

	return r
}

func coffeRoutes(r *chi.Mux) {
	r.Post("/coffes", coffes.Create)
	r.Put("/coffes/{id}", coffes.Update)
	r.Delete("/coffes/{id}", coffes.Delete)
	r.Get("/coffes", coffes.List)
	r.Get("/coffes/{id}", coffes.Get)
}

func cartRoutes(r *chi.Mux) {
	r.Get("/cart", cart.List)
	r.Post("/cart", cart.Create)
	r.Delete("/cart/{id}", cart.Delete)
	r.Put("/cart/{id}", cart.UpdateBy)
	r.Get("/cart/{id}", cart.Get)
}
