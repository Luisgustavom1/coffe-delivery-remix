package cart

import (
	"coffee-delivery-remix/api/models"
	"coffee-delivery-remix/api/pkg/serialize"
	"coffee-delivery-remix/api/services/db"
	"log"
)

func GetAll() ([]models.Cart, error) {
	cart := []models.Cart{}
	connection, err := db.OpenConnection()
	if err != nil {
		return cart, err
	}
	defer connection.Close()

	sql := `SELECT cart.id, cart.quantity, json_build_object('id', coffes.id, 'img', coffes.img, 'price', coffes.price, 'title', coffes.title, 'description', coffes.description, 'stok', coffes.stok, 'categories', coffes.categories) FROM cart JOIN coffes ON cart.productId = coffes.id`

	rows, err := connection.Query(sql)
	if err != nil {
		return cart, err
	}

	for rows.Next() {
		var Cart models.Cart
		var jsonProduct []byte

		err = rows.Scan(&Cart.ID, &Cart.Quantity, &jsonProduct)
		if err != nil {
			log.Printf("Error: %v\n", err.Error())
			continue
		}
		serialize.Cart(jsonProduct, &Cart)

		cart = append(cart, Cart)
	}

	return cart, err
}
