package repository

import (
	"coffee-delivery-remix/api/entities"
	"encoding/json"
	"log"
)

func (p *ProductRepository) UpdateBy(id int64, product entities.Product) (int64, error) {
	categoriesSerialized, err := json.Marshal(product.Categories)
	if err != nil {
		log.Printf("Erro ao fazer o decode das categories: %v", err)
		return id, err
	}

	res, err := p.Conn.Exec("UPDATE product SET img=$2, price=$3, title=$4, description=$5, stok=$6, categories=$7, type=$8 WHERE id=$1",
		id,
		product.Img,
		product.Price,
		product.Title,
		product.Description,
		product.Stok,
		categoriesSerialized,
		product.Type)
	if err != nil {
		return id, err
	}

	return res.RowsAffected()
}
