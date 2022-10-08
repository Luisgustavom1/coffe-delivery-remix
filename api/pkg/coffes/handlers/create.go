package coffes

import (
	"coffe-delivery-remix/api/entities"
	coffes "coffe-delivery-remix/api/pkg/coffes/models"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func Create(w http.ResponseWriter, request *http.Request) {
	var coffe entities.Coffe[[]string]

	err := json.NewDecoder(request.Body).Decode(&coffe)
	if err != nil {
		log.Printf("Erro ao fazer o decode do json: %w", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	id, err := coffes.Insert(coffe)

	var response map[string]any

	if err != nil {
		response = map[string]any{
			"Message": fmt.Sprintf("Ocorreu um erro ao tentar inserir: %v", err),
		}
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		response = map[string]any{
			"Message": fmt.Sprintf("Café inserido com sucesso! ID: %d", id),
		}
		w.WriteHeader(http.StatusCreated)
	}

	w.Header().Add("Content-type", "application/json")
	json.NewEncoder(w).Encode(response)
}
