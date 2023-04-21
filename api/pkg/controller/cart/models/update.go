package cart

import (
	"coffee-delivery-remix/api/entities"
	"coffee-delivery-remix/api/services/db"
)

func UpdateProductByCartId(id int64, cartUpdate entities.CartProductSimple) (int64, error) {
	connection, err := db.OpenConnection()
	if err != nil {
		return 0, err
	}
	defer connection.Close()

	sql := `UPDATE cart_product SET quantity=$3 WHERE cart_id=$1 AND product_id=$2`

	row, err := connection.Exec(sql, id, cartUpdate.ProductId, cartUpdate.Quantity)
	if err != nil {
		return 0, err
	}

	return row.RowsAffected()
}
