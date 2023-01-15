package cart

import "coffe-delivery-remix/rest/infrastructure/db"

func DeleteBy(id int64) (int64, error) {
	connection, err := db.OpenConnection()
	if err != nil {
		return 0, err
	}
	defer connection.Close()

	sql := `DELETE FROM cart WHERE id=$1`

	row, err := connection.Exec(sql, id)
	if err != nil {
		return 0, err
	}

	return row.RowsAffected()
}
