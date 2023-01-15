package coffes

import (
	"coffe-delivery-remix/rest/api/presenter"
	"coffe-delivery-remix/rest/infra/db"
	"encoding/json"
	"log"
)

func Insert(coffe presenter.Coffe) (id int64, err error) {
	connection, err := db.OpenConnection()
	if err != nil {
		return
	}
	defer connection.Close()

	sql := `INSERT INTO coffes (img, price, title, description, stok, categories) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id`

	categoriesAdapted, err := json.Marshal(coffe.Categories)
	if err != nil {
		log.Printf("Erro ao fazer o decode do json: %v", err)
		return
	}

	err = connection.QueryRow(sql, coffe.Img, coffe.Price, coffe.Title, coffe.Description, coffe.Stok, categoriesAdapted).Scan(&id)

	return id, err
}
