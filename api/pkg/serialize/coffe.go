package serialize

import (
	"coffee-delivery-remix/api/models"
	"encoding/json"
)

func Coffe(coffe models.CoffeSimple, coffeSerialized *models.Coffe) error {
	var coffeCategoriesAsArray []string

	err := json.Unmarshal([]byte(coffe.Categories), &coffeCategoriesAsArray)

	*coffeSerialized = models.Coffe{
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
