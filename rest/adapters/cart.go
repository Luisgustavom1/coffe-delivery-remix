package adapters

import (
	"coffe-delivery-remix/rest/api/presenter"
	"coffe-delivery-remix/rest/entity"
	"encoding/json"
)

func Cart(jsonProduct []byte, cart entity.ProductCart, cartAdapted *presenter.ProductCart) (err error) {
	var coffe entity.Coffe
	json.Unmarshal([]byte(jsonProduct), &coffe)

	*cartAdapted = presenter.ProductCart{
		ID:       cart.ID,
		Quantity: cart.Quantity,
	}

	err = Coffe(coffe, &cartAdapted.Product)

	return err
}
