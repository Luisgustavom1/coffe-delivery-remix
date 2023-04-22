package products

import (
	repository "coffee-delivery-remix/api/infra/repository/products"
)

type ProductUseCase struct {
	productRepository repository.ProductRepository
}

func NewProductUseCase(repo repository.ProductRepository) *ProductUseCase {
	return &ProductUseCase{
		productRepository: repo,
	}
}
