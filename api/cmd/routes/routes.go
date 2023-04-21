package routes

import (
	cart "coffee-delivery-remix/api/pkg/controller/cart/handlers"
	checkout "coffee-delivery-remix/api/pkg/controller/checkout/handlers"
	products "coffee-delivery-remix/api/pkg/controller/products/handlers"

	"github.com/go-chi/chi/v5"
)

func Routes() *chi.Mux {
	r := chi.NewRouter()

	productsRoutes(r)
	cartRoutes(r)
	checkoutRoutes(r)

	return r
}

func productsRoutes(r *chi.Mux) {
	r.Post("/products", products.Create)
	r.Put("/products/{id}", products.Update)
	r.Delete("/products/{id}", products.Delete)
	r.Get("/products", products.List)
	r.Get("/products/{id}", products.Get)
}

func cartRoutes(r *chi.Mux) {
	r.Get("/cart", cart.List)
	r.Post("/cart", cart.Create)
	r.Delete("/cart/{id}", cart.Delete)
	r.Put("/cart/{id}", cart.UpdateBy)
	r.Get("/cart/{id}", cart.Get)
}

func checkoutRoutes(r *chi.Mux) {
	r.Post("/checkout", checkout.Checkout)
}
