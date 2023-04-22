package cart

import (
	cart "coffee-delivery-remix/api/pkg/controller/cart/models"
	products "coffee-delivery-remix/api/pkg/controller/products/models"
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
