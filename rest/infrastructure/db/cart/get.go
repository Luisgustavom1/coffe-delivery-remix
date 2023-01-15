package cart

import (
	"coffe-delivery-remix/rest/adapters"
	"coffe-delivery-remix/rest/api/presenter"
	"coffe-delivery-remix/rest/entity"
	"coffe-delivery-remix/rest/infrastructure/db"
)

func GetById(id int64) (cartAdapted presenter.ProductCart, err error) {
	connection, err := db.OpenConnection()
	if err != nil {
		return
	}
	defer connection.Close()

	sql := `SELECT cart.id, cart.quantity, json_build_object('id', coffes.id, 'img', coffes.img, 'price', coffes.price, 'title', coffes.title, 'description', coffes.description, 'stok', coffes.stok, 'categories', coffes.categories) FROM cart JOIN coffes ON cart.productId = coffes.id WHERE cart.id=$1`

	row := connection.QueryRow(sql, id)

	var cart entity.ProductCart
	var jsonProduct []byte
	err = row.Scan(&cart.ID, &cart.Quantity, &jsonProduct)

	err = adapters.Cart(jsonProduct, cart, &cartAdapted)

	return cartAdapted, err
}

func GetByProductId(id int64) (cartAdapted presenter.ProductCart, err error) {
	connection, err := db.OpenConnection()
	if err != nil {
		return
	}
	defer connection.Close()

	sql := `SELECT cart.id, cart.quantity, json_build_object('id', coffes.id, 'img', coffes.img, 'price', coffes.price, 'title', coffes.title, 'description', coffes.description, 'stok', coffes.stok, 'categories', coffes.categories) FROM cart JOIN coffes ON cart.productId = coffes.id WHERE product.id=$1`

	row := connection.QueryRow(sql, id)

	var cart entity.ProductCart
	var jsonProduct []byte
	err = row.Scan(&cart.ID, &cart.Quantity, &jsonProduct)

	err = adapters.Cart(jsonProduct, cart, &cartAdapted)

	return cartAdapted, err
}
