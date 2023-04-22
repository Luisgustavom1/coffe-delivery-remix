package cart

import (
	cart "coffee-delivery-remix/api/infra/repository/cart"
	products "coffee-delivery-remix/api/infra/repository/products"
)

type CartUseCase struct {
	cartRepository    cart.CartRepository
	productRepository products.ProductRepository
}

func NewCartUseCase(c cart.CartRepository, p products.ProductRepository) *CartUseCase {
	return &CartUseCase{
		cartRepository:    c,
		productRepository: p,
	}
}
