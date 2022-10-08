package coffes

import (
	"coffe-delivery-remix/api/entities"
	"coffe-delivery-remix/api/pkg/coffes/models"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

func Update(w http.ResponseWriter, request *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(request, "id"))
	if err != nil {
		log.Printf("Erro ao fazer o parse do query param id: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	var coffe entities.Coffe[[]string]

	err = json.NewDecoder(request.Body).Decode(&coffe)
	if err != nil {
		log.Printf("Erro ao fazer o decode do json: %w", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	rowsAffected, err := coffes.UpdateBy(int64(id), coffe)
	if err != nil {
		log.Printf("Erro ao atualizar registro: %v", id)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	if rowsAffected > 1 {
		log.Printf("Error: foram atualizadas %d registros", rowsAffected)
	}

	response := map[string]any{
		"Error":   false,
		"Message": "dados foram atualizados com sucesso!",
	}

	w.Header().Add("Content-type", "application/json")
	json.NewEncoder(w).Encode(response)
}