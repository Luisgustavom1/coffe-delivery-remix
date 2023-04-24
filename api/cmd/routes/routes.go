package routes

import (
	repository_cart "coffee-delivery-remix/api/infra/repository/cart"
	repository_product "coffee-delivery-remix/api/infra/repository/products"
	
	"coffee-delivery-remix/api/usecase/cart"
	"coffee-delivery-remix/api/usecase/products"
	"coffee-delivery-remix/api/usecase/checkout"
	"database/sql"

	"github.com/go-chi/chi/v5"
)

func Routes(conn *sql.DB) *chi.Mux {
	r := chi.NewRouter()

	cartRepository := *repository_cart.NewCartRepository(conn)
	productRepository := *repository_product.NewProductRepository(conn)

	productUsecase := *products.NewProductUseCase(productRepository)
	cartUsecase := *cart.NewCartUseCase(cartRepository, productRepository)
	checkoutUsecase := *checkout.NewCheckoutUseCase(cartRepository)

	productsRoutes(r, productUsecase)
	cartRoutes(r, cartUsecase)
	checkoutRoutes(r, checkoutUsecase)

	return r
}

func productsRoutes(r *chi.Mux, p products.ProductUseCase) {
	r.Post("/products", p.Create)
	r.Put("/products/{id}", p.Update)
	r.Delete("/products/{id}", p.Delete)
	r.Get("/products", p.List)
	r.Get("/products/{id}", p.Get)
}

func cartRoutes(r *chi.Mux, c cart.CartUseCase) {
	r.Post("/cart", c.Create)
	r.Delete("/cart/{id}", c.DeleteCart)
	r.Get("/cart/{id}", c.Get)
	
	r.Put("/cart/{id}/product/{productId}", c.UpdateCartProduct)
	r.Delete("/cart/{id}/product/{productId}", c.DeleteCartProduct)
	r.Post("/cart/{id}/product", c.CreateCartProduct)
}

func checkoutRoutes(r *chi.Mux, c checkout.CheckoutUseCase) {
	r.Post("/checkout/{id}", c.Checkout)
}
