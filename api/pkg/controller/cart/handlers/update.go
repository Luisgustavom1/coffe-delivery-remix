package cart

import (
	"coffee-delivery-remix/api/models"
	cart "coffee-delivery-remix/api/pkg/controller/cart/models"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

func UpdateBy(w http.ResponseWriter, request *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(request, "id"))
	if err != nil {
		log.Printf("Erro ao fazer o parser do query param id: %v\n", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	var cartUpdate models.CartSimple

	err = json.NewDecoder(request.Body).Decode(&cartUpdate)
	if err != nil {
		log.Printf("Erro ao fazer o decode do json: %v\n", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	rowsAffected, err := cart.UpdateBy(int64(id), cartUpdate)
	if err != nil {
		log.Printf("Erro ao atualizar registro: %v", id)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	if rowsAffected > 1 {
		log.Printf("Error: foram atualizadas %d registros", rowsAffected)
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]any{
		"Message": "dados foram atualizados com sucesso!",
	})
}
