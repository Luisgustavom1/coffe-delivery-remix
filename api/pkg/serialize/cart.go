package serialize

import (
	"coffe-delivery-remix/api/models"
	"encoding/json"
)

func Cart(jsonProduct []byte, cartSerialized *models.Cart) (err error) {
	var coffe models.CoffeSimple
	json.Unmarshal([]byte(jsonProduct), &coffe)

	err = Coffe(coffe, &cartSerialized.Product)

	return err
}
