package cart

import "coffee-delivery-remix/api/services/db"

func DeleteById(id int64) (int64, error) {
	connection, err := db.OpenConnection()
	if err != nil {
		return 0, err
	}
	defer connection.Close()

	sql := `DELETE FROM cart WHERE id=$1`

	row, err := connection.Exec(sql, id)
	if err != nil {
		return 0, err
	}

	return row.RowsAffected()
}

func DeleteCartProductById(cartId int64, productId int64) (int64, error) {
	connection, err := db.OpenConnection()
	if err != nil {
		return 0, err
	}
	defer connection.Close()

	sql := `DELETE FROM cart_product WHERE product_id=$1 AND cart_id=$2`

	row, err := connection.Exec(sql, productId, cartId)
	if err != nil {
		return 0, err
	}

	return row.RowsAffected()
}