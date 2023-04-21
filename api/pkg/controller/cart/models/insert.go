package cart

import (
	"coffee-delivery-remix/api/entities"
	"coffee-delivery-remix/api/services/db"
)

func Insert(cart entities.CartSimple) (id int64, err error) {
	connection, err := db.OpenConnection()
	if err != nil {
		return
	}
	defer connection.Close()

	sql := `INSERT INTO cart (quantity, productId) VALUES ($1, $2) RETURNING id`

	err = connection.QueryRow(sql, cart.Quantity, cart.ProductId).Scan(&id)

	return id, err
}
