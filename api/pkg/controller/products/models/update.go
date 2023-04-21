package products

import (
	"coffee-delivery-remix/api/entities"
	"coffee-delivery-remix/api/services/db"
)

func UpdateBy(id int64, product entities.Product) (int64, error) {
	connection, err := db.OpenConnection()
	if err != nil {
		return 0, err
	}
	defer connection.Close()

	res, err := connection.Exec("UPDATE product SET img=$2, price=$3, title=$4, description=$5, stok=$6, categories=$7, type=$8 WHERE id=$1", id, product.Img, product.Price, product.Title, product.Description, product.Stok, product.Categories, product.Type)
	if err != nil {
		return 0, err
	}

	return res.RowsAffected()
}
