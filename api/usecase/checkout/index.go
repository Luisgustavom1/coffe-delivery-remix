package checkout

import (
	cart "coffee-delivery-remix/api/infra/repository/cart"
)

type CheckoutUseCase struct {
	cartRepository cart.CartRepository
}

func NewCheckoutUseCase(repo cart.CartRepository) *CheckoutUseCase {
	return &CheckoutUseCase{
		cartRepository: repo,
	}
}
