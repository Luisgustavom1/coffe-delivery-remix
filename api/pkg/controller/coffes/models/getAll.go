package coffes

import (
	"coffe-delivery-remix/api/entities"
	"coffe-delivery-remix/api/pkg/serialize"
	"coffe-delivery-remix/api/services/db"
	"log"
)

func GetAll() ([]entities.Coffe, error) {
	coffes := []entities.Coffe{}
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
		var coffe entities.CoffeSimple
		var coffeSerialized entities.Coffe

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
