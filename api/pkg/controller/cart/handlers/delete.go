package cart

import (
	cart "coffee-delivery-remix/api/pkg/controller/cart/models"
	http_error "coffee-delivery-remix/api/pkg/controller/errors"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

func DeleteCart(w http.ResponseWriter, request *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(request, "id"))
	if err != nil {
		log.Printf("Erro ao fazer o parse do query param id: %v", err)
		http_error.HttpError(w, 500)
		return
	}

	rowsAffected, err := cart.DeleteById(int64(id))
	if err != nil {
		log.Printf("Erro ao deletar registro: %v", id)
		http_error.HttpError(w, 500)
		return
	}

	if rowsAffected > 1 {
		log.Printf("Error: foram deletados %d registros", rowsAffected)
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]any{
		"Message": "Registro deletado com sucesso!",
	})
}

// Turn delete methods more generic
func DeleteCartProduct(w http.ResponseWriter, request *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(request, "id"))
	productId, err := strconv.Atoi(chi.URLParam(request, "productId"))
	if err != nil {
		log.Printf("Erro ao fazer o parse dos query params: %v", err)
		http_error.HttpError(w, 500)
		return
	}

	rowsAffected, err := cart.DeleteCartProductById(int64(id), int64(productId))
	if err != nil {
		log.Printf("Erro ao deletar registro: %v", id)
		http_error.HttpError(w, 500)
		return
	}

	if rowsAffected > 1 {
		log.Printf("Error: foram deletados %d registros", rowsAffected)
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]any{
		"Message": "Produto removido com sucesso!",
	})
}
