package cart

import (
	"coffe-delivery-remix/api/entities"
	"coffe-delivery-remix/api/services/db"
)

func UpdateBy(id int64, cartUpdate entities.ProductCartSimple) (int64, error) {
	connection, err := db.OpenConnection()
	if err != nil {
		return 0, err
	}
	defer connection.Close()

	sql := `UPDATE cart SET quantity=$2 WHERE id=$1`

	row, err := connection.Exec(sql, id, cartUpdate.Quantity)
	if err != nil {
		return 0, err
	}

	return row.RowsAffected()
}
