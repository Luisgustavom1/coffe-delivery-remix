package cart

import (
	cart "coffee-delivery-remix/api/pkg/controller/cart/models"
)

type CartUseCase struct {
	cartRepository cart.CartRepository
}

func NewCartUseCase(repo cart.CartRepository) *CartUseCase {
	return &CartUseCase{
		cartRepository: repo,
	}
}
