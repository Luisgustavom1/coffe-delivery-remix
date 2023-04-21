package cart

import (
	"coffe-delivery-remix/api/models"
	cart "coffe-delivery-remix/api/pkg/controller/cart/models"
	coffes "coffe-delivery-remix/api/pkg/controller/coffes/models"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func Create(w http.ResponseWriter, request *http.Request) {
	var Cart models.CartSimple

	err := json.NewDecoder(request.Body).Decode(&Cart)
	if err != nil {
		log.Printf("Error ao decodificar o json: %v", err)
		return
	}

	w.Header().Add("Content-Type", "application/json")

	var response map[string]any
	coffe, _ := coffes.GetById(Cart.ProductId)
	if coffe.ID == 0 {
		response = map[string]any{
			"Message": "Café não encontrado",
		}
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(response)
		return
	}

	productAlreadyExistsInCart, _ := cart.GetByProductId(Cart.ProductId)
	if productAlreadyExistsInCart.ID != 0 {
		response = map[string]any{
			"Message": "Produto ja existente no carrinho",
		}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response)
		return
	}

	id, err := cart.Insert(Cart)

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
