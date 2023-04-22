package products

import (
	"coffee-delivery-remix/api/entities"
	"encoding/json"
	"log"
)

func (p *ProductRepository) Insert(product entities.Product) (id int64, err error) {
	sql := `INSERT INTO product (img, price, title, description, stok, categories, type) VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id`

	categoriesSerialized, err := json.Marshal(product.Categories)
	if err != nil {
		log.Printf("Erro ao fazer o decode das categories: %v", err)
		return
	}

	err = p.Conn.QueryRow(sql, product.Img, product.Price, product.Title, product.Description, product.Stok, categoriesSerialized, product.Type).Scan(&id)

	return id, err
}
