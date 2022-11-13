package cart

import (
	"coffe-delivery-remix/api/entities"
	"coffe-delivery-remix/api/services/db"
	"encoding/json"
	"log"
)

func GetAll() (cart []entities.CartProduct, err error) {
	connection, err := db.OpenConnection()
	if err != nil {
		return
	}
	defer connection.Close()

	sql := `SELECT cart.id, cart.quantity, json_build_object('id', coffes.id, 'img', coffes.img, 'price', coffes.price, 'title', coffes.title, 'description', coffes.description, 'stok', coffes.stok, 'categories', coffes.categories) FROM cart JOIN coffes ON cart.productId = coffes.id`

	rows, err := connection.Query(sql)
	if err != nil {
		return
	}

	for rows.Next() {
		var productCart entities.CartProduct
		var jsonObject []byte

		err = rows.Scan(&productCart.ID, &productCart.Quantity, &jsonObject)
		if err != nil {
			log.Printf("Error: %v\n", err.Error())
			continue
		}
		json.Unmarshal([]byte(jsonObject), &productCart.Product)

		cart = append(cart, productCart)
	}

	return cart, err
}
