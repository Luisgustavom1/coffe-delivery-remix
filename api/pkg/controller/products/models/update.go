package products

import (
	"coffee-delivery-remix/api/entities"
	"coffee-delivery-remix/api/services/db"
	"encoding/json"
	"log"
)

func UpdateBy(id int64, product entities.Product) (int64, error) {
	connection, err := db.OpenConnection()
	if err != nil {
		return id, err
	}
	defer connection.Close()

	categoriesSerialized, err := json.Marshal(product.Categories)
	if err != nil {
		log.Printf("Erro ao fazer o decode das categories: %v", err)
		return id, err
	}

	res, err := connection.Exec("UPDATE product SET img=$2, price=$3, title=$4, description=$5, stok=$6, categories=$7, type=$8 WHERE id=$1",
		id,
		product.Img,
		product.Price,
		product.Title,
		product.Description,
		product.Stok,
		categoriesSerialized,
		product.Type)
	if err != nil {
		return id, err
	}

	return res.RowsAffected()
}
