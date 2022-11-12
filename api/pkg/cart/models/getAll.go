package cart

import (
	"coffe-delivery-remix/api/entities"
	"coffe-delivery-remix/api/services/db"
	"log"
)

func GetAll() (cart []entities.CartProduct, err error) {
	connection, err := db.OpenConnection()
	if err != nil {
		return
	}
	defer connection.Close()

	sql := "SELECT * FROM cart"

	rows, err := connection.Query(sql)
	if err != nil {
		return
	}

	for rows.Next() {
		var productCart entities.CartProduct

		err = rows.Scan(&productCart.Product, &productCart.Quantity)
		if err != nil {
			log.Printf("Error: %v\n", err.Error())
			continue
		}

		cart = append(cart, productCart)
	}

	return cart, err
}