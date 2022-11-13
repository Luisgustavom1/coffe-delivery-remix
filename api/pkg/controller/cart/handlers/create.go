package cart

import (
	"coffe-delivery-remix/api/entities"
	cart "coffe-delivery-remix/api/pkg/controller/cart/models"
	coffes "coffe-delivery-remix/api/pkg/controller/coffes/models"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func Create(w http.ResponseWriter, request *http.Request) {
	var ProductCart entities.ProductCartSimple

	err := json.NewDecoder(request.Body).Decode(&ProductCart)
	if err != nil {
		log.Printf("Error ao decodificar o json: %v", err)
		return
	}

	w.Header().Add("Content-Type", "application/json")

	var response map[string]any
	coffe, _ := coffes.GetBy(ProductCart.ProductId)
	if coffe.ID == 0 {
		response = map[string]any{
			"Message": "Café não encontrado",
		}
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(response)
		return
	}

	id, err := cart.Insert(ProductCart)

	if err != nil {
		response = map[string]any{
			"Message": fmt.Sprintf("Ocorreu um erro ao inserir o produto no carrinho - %v", err),
		}
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		response = map[string]any{
			"Message": fmt.Sprintf("Produto inserido no carrinho com sucesso! ID: %d", id),
		}
		w.WriteHeader(http.StatusCreated)
	}

	json.NewEncoder(w).Encode(response)
}
