package products

import (
	"coffee-delivery-remix/api/entities"
	products "coffee-delivery-remix/api/pkg/controller/products/models"
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

	var product entities.Product

	err = json.NewDecoder(request.Body).Decode(&product)
	if err != nil {
		log.Printf("Erro ao fazer o decode do json: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	rowsAffected, err := products.UpdateBy(int64(id), product)
	if err != nil {
		log.Printf("Erro ao atualizar registro: %v", id)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	if rowsAffected > 1 {
		log.Printf("Error: foram atualizadas %d registros", rowsAffected)
	}

	response := map[string]any{
		"Message": "dados foram atualizados com sucesso!",
	}

	w.Header().Add("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
