package models

import (
	"coffe-delivery-remix/api/services/db"
)

func Insert(coffe Coffe) (id int64, err error) {
	connection, err := db.OpenConnection()
	if err != nil {
		return
	}
	defer connection.Close()

	sql := `INSERT INTO coffes (img, price, title, description, stok) VALUES ($1, $2, $3, $4, $5) RETURNING id`

	err = connection.QueryRow(sql, coffe.Img, coffe.Price, coffe.Title, coffe.Description, coffe.Stok).Scan(&id)

	return id, err
}
