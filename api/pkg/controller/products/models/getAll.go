package products

import (
	"coffee-delivery-remix/api/entities"
	"coffee-delivery-remix/api/pkg/serialize"
	"coffee-delivery-remix/api/services/db"
	"log"
)

func GetAll() ([]entities.Product, error) {
	products := []entities.Product{}
	connection, err := db.OpenConnection()
	if err != nil {
		return products, err
	}
	defer connection.Close()

	sql := `SELECT * FROM product`

	rows, err := connection.Query(sql)
	if err != nil {
		return products, err
	}

	for rows.Next() {
		var product entities.ProductSimple
		var productSerialized entities.Product

		err = rows.Scan(&product.ID, &product.Img, &product.Price, &product.Title, &product.Description, &product.Stok, &product.Categories)
		if err != nil {
			log.Printf("ERRO: %s\n", err.Error())
			continue
		}

		err = serialize.Product(product, &productSerialized)

		products = append(products, productSerialized)
	}

	return products, err
}
