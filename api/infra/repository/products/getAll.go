package repository

import (
	"coffee-delivery-remix/api/entities"
	"coffee-delivery-remix/api/infra/adapters"
	"log"
)

func (p *ProductRepository) GetAll() ([]entities.Product, error) {
	products := []entities.Product{}

	sql := `SELECT * FROM product`

	rows, err := p.Conn.Query(sql)
	if err != nil {
		return products, err
	}

	for rows.Next() {
		var product entities.ProductSimple
		var productSerialized entities.Product

		err = rows.Scan(
			&product.ID,
			&product.Img,
			&product.Price,
			&product.Title,
			&product.Description,
			&product.Stok,
			&product.Categories,
			&product.Type,
			&product.CreatedAt,
			&product.UpdatedAt,
		)
		if err != nil {
			log.Printf("ERRO: %s\n", err.Error())
			continue
		}

		err = adapters.Product(product, &productSerialized)

		products = append(products, productSerialized)
	}

	return products, err
}
