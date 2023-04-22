package routes

import (
	cart_usecase "coffee-delivery-remix/api/pkg/controller/cart/handlers"
	cart_repo "coffee-delivery-remix/api/pkg/controller/cart/models"
	checkout "coffee-delivery-remix/api/pkg/controller/checkout/handlers"
	products "coffee-delivery-remix/api/pkg/controller/products/handlers"
	"database/sql"

	"github.com/go-chi/chi/v5"
)

func Routes(conn *sql.DB) *chi.Mux {
	r := chi.NewRouter()

	cartRepository := cart_repo.NewCartRepository(conn)
	cartUsecase := *cart_usecase.NewCartUseCase(*cartRepository)
	checkoutUsecase := *checkout.NewCheckoutUseCase(*cartRepository)

	productsRoutes(r)
	cartRoutes(r, cartUsecase)
	checkoutRoutes(r, checkoutUsecase)

	return r
}

func productsRoutes(r *chi.Mux) {
	r.Post("/products", products.Create)
	r.Put("/products/{id}", products.Update)
	r.Delete("/products/{id}", products.Delete)
	r.Get("/products", products.List)
	r.Get("/products/{id}", products.Get)
}

func cartRoutes(r *chi.Mux, c cart_usecase.CartUseCase) {
	r.Post("/cart", c.Create)
	r.Delete("/cart/{id}", c.DeleteCart)
	r.Put("/cart/{id}", c.UpdateBy)
	r.Get("/cart/{id}", c.Get)
	r.Delete("/cart/{id}/product/{productId}", c.DeleteCartProduct)
}

func checkoutRoutes(r *chi.Mux, c checkout.CheckoutUseCase) {
	r.Post("/checkout/{id}", c.Checkout)
}
