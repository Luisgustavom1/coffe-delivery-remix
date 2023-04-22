package repository

import "database/sql"

type ProductRepository struct {
	Conn *sql.DB
}

func NewProductRepository(conn *sql.DB) *ProductRepository {
	return &ProductRepository{conn}
}
