package cart

import (
	"coffe-delivery-remix/api/entities"
	"coffe-delivery-remix/api/services/db"
)

func Insert(cart entities.ProductCartSimple) (id int64, err error) {
	connection, err := db.OpenConnection()
	if err != nil {
		return
	}
	defer connection.Close()

	sql := `INSERT INTO cart (quantity, productId) VALUES ($1, $2) RETURNING id`

	err = connection.QueryRow(sql, cart.Quantity, cart.ProductId).Scan(&id)

	return id, err
}
