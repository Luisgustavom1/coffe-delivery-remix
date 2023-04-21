package products

import (
	"coffee-delivery-remix/api/entities"
	"coffee-delivery-remix/api/pkg/serialize"
	"coffee-delivery-remix/api/services/db"
)

func GetById(id int64) (productSerialized entities.Product, err error) {
	connection, err := db.OpenConnection()
	if err != nil {
		return
	}
	defer connection.Close()

	sql := `SELECT * FROM product WHERE id=$1`
	row := connection.QueryRow(sql, id)
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

	err = serialize.Product(product, &productSerialized)

	return productSerialized, err
}
