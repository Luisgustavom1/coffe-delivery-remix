package adapters

import (
	"coffe-delivery-remix/rest/api/presenter"
	"coffe-delivery-remix/rest/entity"
	"encoding/json"
)

func Coffe(coffe entity.Coffe, coffeAdapted *presenter.Coffe) error {
	var coffeCategoriesAsArray []string

	err := json.Unmarshal([]byte(coffe.Categories), &coffeCategoriesAsArray)

	*coffeAdapted = presenter.Coffe{
		ID:          coffe.ID,
		Img:         coffe.Img,
		Price:       coffe.Price,
		Title:       coffe.Title,
		Description: coffe.Description,
		Stok:        coffe.Stok,
		Categories:  coffeCategoriesAsArray,
	}

	return err
}
