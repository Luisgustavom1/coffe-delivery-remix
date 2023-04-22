package products

import (
	"coffee-delivery-remix/api/entities"
	"coffee-delivery-remix/api/pkg/serialize"
)

func (p *ProductRepository) GetById(id int64) (productSerialized entities.Product, err error) {
	sql := `SELECT * FROM product WHERE id=$1`

	row := p.Conn.QueryRow(sql, id)
	var product entities.ProductSimple

	err = row.Scan(
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
		return productSerialized, err
	}

	err = serialize.Product(product, &productSerialized)

	return productSerialized, err
}
