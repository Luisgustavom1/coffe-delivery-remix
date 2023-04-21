package cart

import (
	"coffee-delivery-remix/api/entities"
	"coffee-delivery-remix/api/pkg/serialize"
	"coffee-delivery-remix/api/services/db"
	"log"
)

func GetById(id int64) (cart entities.Cart, err error) {
	connection, err := db.OpenConnection()
	if err != nil {
		return
	}
	defer connection.Close()

	sql := `
		SELECT 
			cp.cart_id, 
			cp.quantity, 
			json_build_object(
				'id', p.id, 
				'img', p.img, 
				'price', p.price, 
				'title', p.title, 
				'description', p.description, 
				'type', p.type,
				'stok', p.stok, 
				'categories', p.categories
			) 
		FROM cart 
			JOIN cart_product cp ON cp.cart_id = cart.id 
			JOIN product p ON cp.product_id = p.id 
		WHERE cart.id=$1`

	rows, err := connection.Query(sql, id)
	if err != nil {
		return cart, err
	}

	for rows.Next() {
		var cartProduct entities.CartProduct
		var jsonProduct []byte

		err = rows.Scan(&cart.ID, &cartProduct.Quantity, &jsonProduct)
		if err != nil {
			log.Printf("Error: %v\n", err.Error())
			continue
		}
		serialize.Cart(jsonProduct, &cartProduct.Product)

		cart.Products = append(cart.Products, cartProduct)
	}

	return cart, err
}
