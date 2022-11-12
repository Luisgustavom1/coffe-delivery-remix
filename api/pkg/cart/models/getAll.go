package cart

import (
	"coffe-delivery-remix/api/entities"
	"coffe-delivery-remix/api/services/db"
	"log"
)

func GetAll() (cart []entities.CartProductSimple, err error) {
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
		var productCart entities.CartProductSimple

		err = rows.Scan(&productCart.ID, &productCart.Quantity, &productCart.ProductId)
		if err != nil {
			log.Printf("Error: %v\n", err.Error())
			continue
		}

		cart = append(cart, productCart)
	}

	return cart, err
}
