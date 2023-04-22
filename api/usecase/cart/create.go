package cart

import (
	"coffee-delivery-remix/api/infra/repository/cart"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func (c *CartUseCase) Create(w http.ResponseWriter, request *http.Request) {
	var input repository.InsertCartInputDTO

	err := json.NewDecoder(request.Body).Decode(&input)
	if err != nil {
		log.Printf("Error ao decodificar o json: %v", err)
		return
	}

	w.Header().Add("Content-Type", "application/json")

	var response map[string]any
	for _, item := range input.Products {
		product, _ := c.productRepository.GetById(item.ProductId)
		if product.ID == 0 {
			response = map[string]any{
				"Message": "Produto não encontrado",
			}
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(response)
			return
		}
	}

	id, err := c.cartRepository.Insert(input)

	if err != nil {
		response = map[string]any{
			"Message": fmt.Sprintf("Ocorreu um erro criar o carrinho - %v", err),
		}
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		response = map[string]any{
			"Message": fmt.Sprintf("Carrinho criado com sucesso! ID: %d", id),
		}
		w.WriteHeader(http.StatusCreated)
	}

	json.NewEncoder(w).Encode(response)
}
