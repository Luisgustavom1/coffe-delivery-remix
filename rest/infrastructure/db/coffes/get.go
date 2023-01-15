package coffes

import (
	"coffe-delivery-remix/rest/adapters"
	"coffe-delivery-remix/rest/api/presenter"
	"coffe-delivery-remix/rest/entity"
	"coffe-delivery-remix/rest/infrastructure/db"
)

func GetById(id int64) (coffeAdapted presenter.Coffe, err error) {
	connection, err := db.OpenConnection()
	if err != nil {
		return
	}
	defer connection.Close()

	sql := `SELECT * FROM coffes WHERE id=$1`
	row := connection.QueryRow(sql, id)
	var coffe entity.Coffe

	err = row.Scan(&coffe.ID, &coffe.Img, &coffe.Price, &coffe.Title, &coffe.Description, &coffe.Stok, &coffe.Categories)

	err = adapters.Coffe(coffe, &coffeAdapted)

	return coffeAdapted, err
}
