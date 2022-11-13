package cart

import (
	"coffe-delivery-remix/api/entities"
	"coffe-delivery-remix/api/services/db"
)

func GetBy(id int64) (cart entities.ProductCartSimple, err error) {
	connection, err := db.OpenConnection()
	if err != nil {
		return
	}
	defer connection.Close()

	sql := `SELECT * FROM cart WHERE id=$1`

	row := connection.QueryRow(sql, id)

	err = row.Scan(&cart.ID, &cart.Quantity, &cart.ProductId)

	return cart, err
}
