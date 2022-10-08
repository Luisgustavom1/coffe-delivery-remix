package coffes

import (
	"coffe-delivery-remix/api/entities"
	"coffe-delivery-remix/api/services/db"
	"encoding/json"
	"log"
)

func Insert(coffe entities.Coffe[[]string]) (id int64, err error) {
	connection, err := db.OpenConnection()
	if err != nil {
		return
	}
	defer connection.Close()

	sql := `INSERT INTO coffes (img, price, title, description, stok, categories) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id`

	if err != nil {
		log.Printf("Erro ao fazer o serialize do categories: %w", err)
	}

	categoriesSerialized, err := json.Marshal(coffe.Categories)
	if err != nil {
		log.Printf("Erro ao fazer o decode do json: %w", err)
		return
	}

	err = connection.QueryRow(sql, coffe.Img, coffe.Price, coffe.Title, coffe.Description, coffe.Stok, categoriesSerialized).Scan(&id)

	return id, err
}
