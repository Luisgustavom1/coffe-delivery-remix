package products

import (
	"coffee-delivery-remix/api/entities"
	"coffee-delivery-remix/api/services/db"
	"encoding/json"
	"log"
)

func Insert(product entities.Product) (id int64, err error) {
	connection, err := db.OpenConnection()
	if err != nil {
		return
	}
	defer connection.Close()

	sql := `INSERT INTO product (img, price, title, description, stok, categories) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id`

	categoriesSerialized, err := json.Marshal(product.Categories)
	if err != nil {
		log.Printf("Erro ao fazer o decode do json: %v", err)
		return
	}

	err = connection.QueryRow(sql, product.Img, product.Price, product.Title, product.Description, product.Stok, categoriesSerialized).Scan(&id)

	return id, err
}
