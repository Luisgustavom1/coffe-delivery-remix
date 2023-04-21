package products

import (
	products "coffee-delivery-remix/api/pkg/controller/products/models"
	"encoding/json"
	"log"
	"net/http"
)

func List(w http.ResponseWriter, request *http.Request) {
	products, err := products.GetAll()
	if err != nil {
		log.Printf("Erro ao obter registros: %v", err)
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(products)
}
