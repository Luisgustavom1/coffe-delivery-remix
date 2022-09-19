package handlers

import (
	"coffe-delivery-remix/api/models"
	"encoding/json"
	"log"
	"net/http"
)

func List(w http.ResponseWriter, request *http.Request) {
	coffes, err := models.GetAll()
	if err != nil {
		log.Printf("Erro ao obter registros: %v", err)
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(coffes)
}