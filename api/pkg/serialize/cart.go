package serialize

import (
	"coffe-delivery-remix/api/entities"
	"encoding/json"
)

func Cart(jsonProduct []byte, cartSerialized *entities.ProductCart) (err error) {
	var coffe entities.CoffeSimple
	json.Unmarshal([]byte(jsonProduct), &coffe)

	err = Coffe(coffe, &cartSerialized.Product)

	return err
}
