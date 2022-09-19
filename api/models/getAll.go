package models

import (
	"coffe-delivery-remix/api/services/db"
	"log"
)

func GetAll() (coffes []Coffe, err error) {
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
		var coffe Coffe

		err = rows.Scan(&coffe.ID, &coffe.Img, &coffe.Price, &coffe.Title, &coffe.Description, &coffe.Stok)
		if err != nil {
			log.Println("ERRO: $s", err.Error())
			continue
		}

		coffes = append(coffes, coffe)
	}

	return coffes, err
}
