package db

import (
	"coffe-delivery-remix/api/configs"

	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func OpenConnection() (*sql.DB, error) {
	cfg := configs.GetDB()

	sc := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		cfg.User,
		cfg.Pass,
		cfg.Host,
		cfg.Port,
		cfg.Database,
	)

	connection, err := sql.Open("postgres", sc)
	if err != nil {
		panic(err)
	}

	err = connection.Ping()

	return connection, err
}
