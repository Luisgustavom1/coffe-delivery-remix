package routes

import (
	cart_usecase "coffee-delivery-remix/api/pkg/controller/cart/handlers"
	cart_repo "coffee-delivery-remix/api/pkg/controller/cart/models"
	checkout "coffee-delivery-remix/api/pkg/controller/checkout/handlers"
	product_usecase "coffee-delivery-remix/api/pkg/controller/products/handlers"
	product_repo "coffee-delivery-remix/api/pkg/controller/products/models"
	"database/sql"

	"github.com/go-chi/chi/v5"
)

func Routes(conn *sql.DB) *chi.Mux {
	r := chi.NewRouter()

	cartRepository := *cart_repo.NewCartRepository(conn)
	productRepository := *product_repo.NewProductRepository(conn)

	productUsecase := *product_usecase.NewProductUseCase(productRepository)
	cartUsecase := *cart_usecase.NewCartUseCase(cartRepository, productRepository)
	checkoutUsecase := *checkout.NewCheckoutUseCase(cartRepository)

	productsRoutes(r, productUsecase)
	cartRoutes(r, cartUsecase)
	checkoutRoutes(r, checkoutUsecase)

	return r
}

func productsRoutes(r *chi.Mux, p product_usecase.ProductUseCase) {
	r.Post("/products", p.Create)
	r.Put("/products/{id}", p.Update)
	r.Delete("/products/{id}", p.Delete)
	r.Get("/products", p.List)
	r.Get("/products/{id}", p.Get)
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
