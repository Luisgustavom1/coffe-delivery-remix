package coffes

import (
	"coffe-delivery-remix/api/entities"
	"coffe-delivery-remix/api/services/db"
	"encoding/json"
	"log"
)

func GetAll() (coffes []entities.Coffe[[]string], err error) {
	connection, err := db.OpenConnection()
	if err != nil {
		return
	}
	defer connection.Close()

	sql := `SELECT * FROM coffes`

	rows, err := connection.Query(sql)
	if err != nil {
		return
	}

	for rows.Next() {
		var coffe entities.Coffe[string]

		err = rows.Scan(&coffe.ID, &coffe.Img, &coffe.Price, &coffe.Title, &coffe.Description, &coffe.Stok, &coffe.Categories)
		if err != nil {
			log.Println("ERRO: $s", err.Error())
			continue
		}

		var categoriesAsArray []string
		err := json.Unmarshal([]byte(coffe.Categories), &categoriesAsArray)
		if err != nil {
			log.Printf("Erro ao fazer o decode do categories: %w", err)
			continue
		}

		coffeFormatted := entities.Coffe[[]string]{
			ID: coffe.ID,
			Img: coffe.Img,
			Price: coffe.Price,
			Title: coffe.Title,
			Description: coffe.Description,
			Stok: coffe.Stok,
			Categories: categoriesAsArray,
		}
		coffes = append(coffes, coffeFormatted)
	}

	return coffes, err
}
