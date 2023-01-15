package cart

import (
	cart "coffe-delivery-remix/rest/infra/db/cart"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

func Delete(w http.ResponseWriter, request *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(request, "id"))
	if err != nil {
		log.Printf("Erro ao fazer o parse do query param id: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	rowsAffected, err := cart.DeleteBy(int64(id))
	if err != nil {
		log.Printf("Erro ao deletar registro: %v", id)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	if rowsAffected > 1 {
		log.Printf("Error: foram deletados %d registros", rowsAffected)
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]any{
		"Message": "registro foi deletado com sucesso!",
	})
}
