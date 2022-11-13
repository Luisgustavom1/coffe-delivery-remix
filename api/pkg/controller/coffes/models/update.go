package coffes

import (
	"coffe-delivery-remix/api/entities"
	"coffe-delivery-remix/api/services/db"
)

func UpdateBy(id int64, coffe entities.Coffe[[]string]) (int64, error) {
	connection, err := db.OpenConnection()
	if err != nil {
		return 0, err
	}
	defer connection.Close()

	res, err := connection.Exec("UPDATE coffes SET img=$2, price=$3, title=$4, description=$5, stok=$6, categories=$7 WHERE id=$1", id, coffe.Img, coffe.Price, coffe.Title, coffe.Description, coffe.Stok, coffe.Categories)
	if err != nil {
		return 0, err
	}

	return res.RowsAffected()
}
