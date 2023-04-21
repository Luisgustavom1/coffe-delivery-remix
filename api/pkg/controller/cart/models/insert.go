package cart

import (
	"coffee-delivery-remix/api/entities"
	"coffee-delivery-remix/api/services/db"
)

type InsertCartInputDTO struct {
	Products []entities.CartProductSimple `json:"products"`
}

func Insert(cart InsertCartInputDTO) (id int64, err error) {
	connection, err := db.OpenConnection()
	if err != nil {
		return
	}
	defer connection.Close()

	sql := `INSERT INTO cart DEFAULT VALUES RETURNING id`

	err = connection.QueryRow(sql).Scan(&id)

	for _, product := range cart.Products {
		sql := `INSERT INTO cart_product (cart_id, product_id, quantity) VALUES ($1, $2, $3)`

		_, err = connection.Exec(sql, id, product.ProductId, product.Quantity)
	}

	return id, err
}
