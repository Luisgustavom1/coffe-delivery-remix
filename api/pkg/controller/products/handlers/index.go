package products

import (
	"coffee-delivery-remix/api/pkg/controller/products/models"
)

type ProductUseCase struct {
	productRepository products.ProductRepository
}

func NewProductUseCase(repo products.ProductRepository) *ProductUseCase {
	return &ProductUseCase{
		productRepository: repo,
	}
}
