package cart

import (
	"coffe-delivery-remix/api/entities"
	"coffe-delivery-remix/api/pkg/serialize"
	"coffe-delivery-remix/api/services/db"
	"log"
)

func GetAll() (cart []entities.ProductCart, err error) {
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
		var productCart entities.ProductCart
		var jsonProduct []byte

		err = rows.Scan(&productCart.ID, &productCart.Quantity, &jsonProduct)
		if err != nil {
			log.Printf("Error: %v\n", err.Error())
			continue
		}
		serialize.Cart(jsonProduct, &productCart)

		cart = append(cart, productCart)
	}

	return cart, err
}
