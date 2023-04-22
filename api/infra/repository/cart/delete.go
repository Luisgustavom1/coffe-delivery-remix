package repository

func (db *CartRepository) DeleteById(id int64) (int64, error) {
	sql := `DELETE FROM cart WHERE id=$1`

	row, err := db.Conn.Exec(sql, id)
	if err != nil {
		return 0, err
	}

	return row.RowsAffected()
}

func (db *CartRepository) DeleteCartProductById(cartId int64, productId int64) (int64, error) {
	sql := `DELETE FROM cart_product WHERE product_id=$1 AND cart_id=$2`

	row, err := db.Conn.Exec(sql, productId, cartId)
	if err != nil {
		return 0, err
	}

	return row.RowsAffected()
}
