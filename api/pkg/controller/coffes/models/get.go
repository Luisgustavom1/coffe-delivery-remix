package coffes

import (
	"coffe-delivery-remix/api/models"
	"coffe-delivery-remix/api/pkg/serialize"
	"coffe-delivery-remix/api/services/db"
)

func GetById(id int64) (coffeSerialized models.Coffe, err error) {
	connection, err := db.OpenConnection()
	if err != nil {
		return
	}
	defer connection.Close()

	sql := `SELECT * FROM coffes WHERE id=$1`
	row := connection.QueryRow(sql, id)
	var coffe models.CoffeSimple

	err = row.Scan(&coffe.ID, &coffe.Img, &coffe.Price, &coffe.Title, &coffe.Description, &coffe.Stok, &coffe.Categories)

	err = serialize.Coffe(coffe, &coffeSerialized)

	return coffeSerialized, err
}
