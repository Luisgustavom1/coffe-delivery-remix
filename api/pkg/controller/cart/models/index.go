package cart

import "database/sql"

type CartRepository struct {
	Conn *sql.DB
}

func NewCartRepository(conn *sql.DB) *CartRepository {
	return &CartRepository{conn}
}