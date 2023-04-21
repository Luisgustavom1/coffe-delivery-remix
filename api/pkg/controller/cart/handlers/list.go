package cart

import (
	cart "coffee-delivery-remix/api/pkg/controller/cart/models"
	"encoding/json"
	"log"
	"net/http"
)

func List(w http.ResponseWriter, request *http.Request) {
	cart, err := cart.GetAll()
	if err != nil {
		log.Printf("Erro ao obter registros: %v", err)
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(cart)
}
