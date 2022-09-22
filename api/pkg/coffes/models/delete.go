package coffes

import "coffe-delivery-remix/api/services/db"

func DeleteBy(id int64) (int64, error) {
	connection, err := db.OpenConnection()
	if err != nil {
		return 0, err
	}
	defer connection.Close()

	res, err := connection.Exec("DELETE FROM coffes WHERE id=$1", id)
	if err != nil {
		return 0, err
	}

	return res.RowsAffected()
}
