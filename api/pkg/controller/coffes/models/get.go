package coffes

import (
	"coffe-delivery-remix/api/services/db"
	"coffe-delivery-remix/api/entities"
)

func GetBy(id int64) (coffe entities.CoffeSimple, err error) {
	connection, err := db.OpenConnection()
	if err != nil {
		return
	}
	defer connection.Close()

	sql := `SELECT * FROM coffes WHERE id=$1`

	row := connection.QueryRow(sql, id)

	err = row.Scan(&coffe.ID, &coffe.Img, &coffe.Price, &coffe.Title, &coffe.Description, &coffe.Stok, &coffe.Categories)

	return coffe, err
}
