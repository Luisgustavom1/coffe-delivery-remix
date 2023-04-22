package cart

import (
	"coffee-delivery-remix/api/entities"
)

func (db *CartRepository) UpdateProductByCartId(id int64, cartUpdate entities.CartProductSimple) (int64, error) {
	sql := `UPDATE cart_product SET quantity=$3 WHERE cart_id=$1 AND product_id=$2`

	row, err := db.Conn.Exec(sql, id, cartUpdate.ProductId, cartUpdate.Quantity)
	if err != nil {
		return 0, err
	}

	return row.RowsAffected()
}
