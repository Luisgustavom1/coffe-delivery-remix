package coffes

import (
	"coffee-delivery-remix/api/models"
	"coffee-delivery-remix/api/pkg/serialize"
	"coffee-delivery-remix/api/services/db"
	"log"
)

func GetAll() ([]models.Coffe, error) {
	coffes := []models.Coffe{}
	connection, err := db.OpenConnection()
	if err != nil {
		return coffes, err
	}
	defer connection.Close()

	sql := `SELECT * FROM coffes`

	rows, err := connection.Query(sql)
	if err != nil {
		return coffes, err
	}

	for rows.Next() {
		var coffe models.CoffeSimple
		var coffeSerialized models.Coffe

		err = rows.Scan(&coffe.ID, &coffe.Img, &coffe.Price, &coffe.Title, &coffe.Description, &coffe.Stok, &coffe.Categories)
		if err != nil {
			log.Printf("ERRO: %s\n", err.Error())
			continue
		}

		err = serialize.Coffe(coffe, &coffeSerialized)

		coffes = append(coffes, coffeSerialized)
	}

	return coffes, err
}
