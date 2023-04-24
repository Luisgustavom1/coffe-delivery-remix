package cart

import (
	"coffee-delivery-remix/api/entities"
	http_error "coffee-delivery-remix/api/infra/errors"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

type Input struct {
	Quantity int64 `json:"quantity"`
}

func (c *CartUseCase) UpdateCartProduct(w http.ResponseWriter, request *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(request, "id"))
	productId, err := strconv.Atoi(chi.URLParam(request, "productId"))
	if err != nil {
		log.Printf("Erro ao fazer o parse dos query params: %v", err)
		http_error.HttpError(w, 500)
		return
	}

	var input Input

	err = json.NewDecoder(request.Body).Decode(&input)
	if err != nil {
		log.Printf("Erro ao fazer o decode do json: %v\n", err)
		http_error.HttpError(w, 500)
		return
	}

	rowsAffected, err := c.cartRepository.UpdateCartProductByCartId(int64(id), entities.CartProductSimple{
		ProductId: int64(productId),
		Quantity:  input.Quantity,
	})
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
