package cart

import (
	"coffee-delivery-remix/api/models"
	"coffee-delivery-remix/api/pkg/serialize"
	"coffee-delivery-remix/api/services/db"
)

func GetById(id int64) (cart models.Cart, err error) {
	connection, err := db.OpenConnection()
	if err != nil {
		return
	}
	defer connection.Close()

	sql := `SELECT cart.id, cart.quantity, json_build_object('id', coffes.id, 'img', coffes.img, 'price', coffes.price, 'title', coffes.title, 'description', coffes.description, 'stok', coffes.stok, 'categories', coffes.categories) FROM cart JOIN coffes ON cart.productId = coffes.id WHERE cart.id=$1`

	row := connection.QueryRow(sql, id)

	var jsonProduct []byte
	err = row.Scan(&cart.ID, &cart.Quantity, &jsonProduct)

	err = serialize.Cart(jsonProduct, &cart)

	return cart, err
}

func GetByProductId(id int64) (cart models.Cart, err error) {
	connection, err := db.OpenConnection()
	if err != nil {
		return
	}
	defer connection.Close()

	sql := `SELECT cart.id, cart.quantity, json_build_object('id', coffes.id, 'img', coffes.img, 'price', coffes.price, 'title', coffes.title, 'description', coffes.description, 'stok', coffes.stok, 'categories', coffes.categories) FROM cart JOIN coffes ON cart.productId = coffes.id WHERE product.id=$1`

	row := connection.QueryRow(sql, id)

	var jsonProduct []byte
	err = row.Scan(&cart.ID, &cart.Quantity, &jsonProduct)

	err = serialize.Cart(jsonProduct, &cart)

	return cart, err
}
