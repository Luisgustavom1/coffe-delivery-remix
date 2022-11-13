package coffes

import (
	coffes "coffe-delivery-remix/api/pkg/controller/coffes/models"
	"encoding/json"
	"log"
	"net/http"
)

func List(w http.ResponseWriter, request *http.Request) {
	coffes, err := coffes.GetAll()
	if err != nil {
		log.Printf("Erro ao obter registros: %v", err)
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(coffes)
}
