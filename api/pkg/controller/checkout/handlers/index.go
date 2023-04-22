package checkout

import (
	cart "coffee-delivery-remix/api/pkg/controller/cart/models"
)

type CheckoutUseCase struct {
	cartRepository cart.CartRepository
}

func NewCheckoutUseCase(repo cart.CartRepository) *CheckoutUseCase {
	return &CheckoutUseCase{
		cartRepository: repo,
	}
}
