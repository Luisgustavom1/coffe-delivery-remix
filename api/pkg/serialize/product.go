package serialize

import (
	"coffee-delivery-remix/api/entities"
	"encoding/json"
)

func Product(product entities.ProductSimple, productSerialized *entities.Product) error {
	var productCategoriesAsArray []string

	err := json.Unmarshal([]byte(product.Categories), &productCategoriesAsArray)

	*productSerialized = entities.Product{
		ID:          product.ID,
		Img:         product.Img,
		Price:       product.Price,
		Title:       product.Title,
		Description: product.Description,
		Stok:        product.Stok,
		Categories:  productCategoriesAsArray,
	}

	return err
}
