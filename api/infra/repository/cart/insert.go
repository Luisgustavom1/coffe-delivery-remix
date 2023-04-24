package repository

import (
	"coffee-delivery-remix/api/entities"
)

type InsertCartInputDTO struct {
	Products []entities.CartProductSimple `json:"products"`
}

func (db *CartRepository) Insert(cart InsertCartInputDTO) (id int64, err error) {
	sql := `INSERT INTO cart DEFAULT VALUES RETURNING id`

	err = db.Conn.QueryRow(sql).Scan(&id)

	for _, product := range cart.Products {
		sql := `INSERT INTO cart_product (cart_id, product_id, quantity) VALUES ($1, $2, $3)`

		_, err = db.Conn.Exec(sql, id, product.ProductId, product.Quantity)
	}

	return id, err
}

func (db *CartRepository) InsertCartProductByCartId(cartId int64, cart_product entities.CartProductSimple) (id int64, err error) {
	sql := `INSERT INTO cart_product (cart_id, product_id, quantity) VALUES ($1, $2, $3) RETURNING id`

	err = db.Conn.QueryRow(sql, cartId, cart_product.ProductId, cart_product.Quantity).Scan(&id)

	return id, err
}
