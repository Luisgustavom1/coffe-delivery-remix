package cart

import (
	"coffe-delivery-remix/rest/entity"
	cart "coffe-delivery-remix/rest/infra/db/cart"
	coffes "coffe-delivery-remix/rest/infra/db/coffes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func Create(w http.ResponseWriter, request *http.Request) {
	var ProductCart entity.ProductCart

	err := json.NewDecoder(request.Body).Decode(&ProductCart)
	if err != nil {
		log.Printf("Error ao decodificar o json: %v", err)
		return
	}

	w.Header().Add("Content-Type", "application/json")

	var response map[string]any
	coffe, _ := coffes.GetById(ProductCart.ProductId)
	if coffe.ID == 0 {
		response = map[string]any{
			"Message": "Café não encontrado",
		}
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(response)
		return
	}

	productAlreadyExistsInCart, _ := cart.GetByProductId(ProductCart.ProductId)
	if productAlreadyExistsInCart.ID != 0 {
		response = map[string]any{
			"Message": "Produto ja existente no carrinho",
		}
		w.WriteHeader(http.StatusBadRequest)
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
