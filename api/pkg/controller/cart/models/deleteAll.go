package cart

import "coffee-delivery-remix/api/services/db"

func DeleteAll() (int64, error) {
	connection, err := db.OpenConnection()
	if err != nil {
		return 0, err
	}
	defer connection.Close()

	sql := `DELETE FROM cart`

	row, err := connection.Exec(sql)
	if err != nil {
		return 0, err
	}

	return row.RowsAffected()
}
