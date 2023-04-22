package cart

import (
	"coffee-delivery-remix/api/entities"
	http_error "coffee-delivery-remix/api/pkg/controller/errors"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

func (c *CartUseCase) UpdateBy(w http.ResponseWriter, request *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(request, "id"))
	if err != nil {
		log.Printf("Erro ao fazer o parser do query param id: %v\n", err)
		http_error.HttpError(w, 500)
		return
	}

	var cartUpdate entities.CartProductSimple

	err = json.NewDecoder(request.Body).Decode(&cartUpdate)
	if err != nil {
		log.Printf("Erro ao fazer o decode do json: %v\n", err)
		http_error.HttpError(w, 500)
		return
	}

	rowsAffected, err := c.cartRepository.UpdateProductByCartId(int64(id), cartUpdate)
	if err != nil {
		log.Printf("Erro ao atualizar registro: %v", id)
		http_error.HttpError(w, 500)
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
