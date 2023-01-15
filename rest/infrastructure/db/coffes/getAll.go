package coffes

import (
	"coffe-delivery-remix/rest/adapters"
	"coffe-delivery-remix/rest/api/presenter"
	"coffe-delivery-remix/rest/entity"
	"coffe-delivery-remix/rest/infrastructure/db"
	"log"
)

func GetAll() ([]presenter.Coffe, error) {
	coffes := []presenter.Coffe{}
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
		var coffe entity.Coffe
		var coffeAdapted presenter.Coffe

		err = rows.Scan(&coffe.ID, &coffe.Img, &coffe.Price, &coffe.Title, &coffe.Description, &coffe.Stok, &coffe.Categories)
		if err != nil {
			log.Printf("ERRO: %s\n", err.Error())
			continue
		}

		err = adapters.Coffe(coffe, &coffeAdapted)

		coffes = append(coffes, coffeAdapted)
	}

	return coffes, err
}
