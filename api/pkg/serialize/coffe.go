package serialize

import (
	"coffe-delivery-remix/api/entities"
	"encoding/json"
)

func Coffe(coffe entities.CoffeSimple, coffeSerialized *entities.Coffe) error {
	var coffeCategoriesAsArray []string

	err := json.Unmarshal([]byte(coffe.Categories), &coffeCategoriesAsArray)

	*coffeSerialized = entities.Coffe{
		ID: coffe.ID,
		Img: coffe.Img,
		Price: coffe.Price,
		Title: coffe.Title,
		Description: coffe.Description,
		Stok: coffe.Stok,
		Categories: coffeCategoriesAsArray,
	}

	return err
}
