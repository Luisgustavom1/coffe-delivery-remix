package products

import (
	"encoding/json"
	"log"
	"net/http"
)

func (p *ProductUseCase) List(w http.ResponseWriter, request *http.Request) {
	products, err := p.productRepository.GetAll()
	if err != nil {
		log.Printf("Erro ao obter registros: %v", err)
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(products)
}
