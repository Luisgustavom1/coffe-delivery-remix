package cart

import (
	"coffe-delivery-remix/rest/adapters"
	"coffe-delivery-remix/rest/api/presenter"
	"coffe-delivery-remix/rest/entity"
	"coffe-delivery-remix/rest/infrastructure/db"
	"log"
)

func GetAll() ([]entity.ProductCart, error) {
	cart := []entity.ProductCart{}
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
		var productCart entity.ProductCart
		var productCartAdapted presenter.ProductCart
		var jsonProduct []byte

		err = rows.Scan(&productCart.ID, &productCart.Quantity, &jsonProduct)
		if err != nil {
			log.Printf("Error: %v\n", err.Error())
			continue
		}
		adapters.Cart(jsonProduct, productCart, &productCartAdapted)

		cart = append(cart, productCart)
	}

	return cart, err
}
