package models

import "coffe-delivery-remix/api/services/db"

func GetBy(id int64) (coffe Coffe, err error) {
	connection, err := db.OpenConnection()
	if err != nil {
		return
	}
	defer connection.Close()

	sql := `SELECT * FROM coffes WHERE id=$1`

	row := connection.QueryRow(sql, id)

	err = row.Scan(&coffe.ID, &coffe.Img, &coffe.Price, &coffe.Title, &coffe.Description, &coffe.Stok)

	return coffe, err
}
