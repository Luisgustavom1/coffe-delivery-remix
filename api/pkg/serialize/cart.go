package serialize

import (
	"coffee-delivery-remix/api/entities"
	"encoding/json"
)

func Cart(jsonProduct []byte, cartSerialized *entities.Product) (err error) {
	var product entities.ProductSimple
	json.Unmarshal([]byte(jsonProduct), &product)

	err = Product(product, cartSerialized)

	return err
}