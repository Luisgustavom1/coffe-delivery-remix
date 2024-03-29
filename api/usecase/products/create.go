package products

import (
	"coffee-delivery-remix/api/entities"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func (p *ProductUseCase) Create(w http.ResponseWriter, request *http.Request) {
	var product entities.Product

	err := json.NewDecoder(request.Body).Decode(&product)
	if err != nil {
		log.Printf("Erro ao fazer o decode do json: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	id, err := p.productRepository.Insert(product)

	var response map[string]any

	if err != nil {
		response = map[string]any{
			"Message": fmt.Sprintf("Ocorreu um erro ao tentar inserir: %v", err),
		}
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		response = map[string]any{
			"Message": fmt.Sprintf("Produto inserido com sucesso! ID: %d", id),
		}
		w.WriteHeader(http.StatusCreated)
	}

	w.Header().Add("Content-type", "application/json")
	json.NewEncoder(w).Encode(response)
}
