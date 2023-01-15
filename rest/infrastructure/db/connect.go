package db

import (
	"coffe-delivery-remix/rest/configs"

	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func OpenConnection() (*sql.DB, error) {
	cfg := configs.GetDB()

	sc := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.Host,
		cfg.Port,
		cfg.User,
		cfg.Pass,
		cfg.Database,
	)

	connection, err := sql.Open("postgres", sc)
	if err != nil {
		panic(err)
	}

	err = connection.Ping()

	return connection, err
}
