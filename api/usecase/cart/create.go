package cart

import (
	"coffee-delivery-remix/api/entities"
	http_error "coffee-delivery-remix/api/infra/errors"
	repository "coffee-delivery-remix/api/infra/repository/cart"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

func (c *CartUseCase) Create(w http.ResponseWriter, request *http.Request) {
	var input repository.InsertCartInputDTO

	err := json.NewDecoder(request.Body).Decode(&input)
	if err != nil {
		log.Printf("Error ao decodificar o json: %v", err)
		return
	}

	w.Header().Add("Content-Type", "application/json")

	var response map[string]any
	for _, item := range input.Products {
		product, _ := c.productRepository.GetById(item.ProductId)
		if product.ID == 0 {
			response = map[string]any{
				"Message": "Produto não encontrado",
			}
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(response)
			return
		}
	}

	id, err := c.cartRepository.Insert(input)

	if err != nil {
		response = map[string]any{
			"Message": fmt.Sprintf("Ocorreu um erro criar o carrinho - %v", err),
		}
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		response = map[string]any{
			"Message": fmt.Sprintf("Carrinho criado com sucesso! ID: %d", id),
		}
		w.WriteHeader(http.StatusCreated)
	}

	json.NewEncoder(w).Encode(response)
}

func (c *CartUseCase) CreateCartProduct(w http.ResponseWriter, request *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(request, "id"))
	if err != nil {
		log.Printf("Erro ao fazer o parse dos query params: %v", err)
		http_error.HttpError(w, 500)
		return
	}

	var input entities.CartProductSimple
	err = json.NewDecoder(request.Body).Decode(&input)
	if err != nil {
		log.Printf("Error ao decodificar o json: %v", err)
		return
	}

	cartProduct, _ := c.cartRepository.GetCartProductById(int64(id), input.ProductId)
	if cartProduct.ProductId == 0 {
		w.WriteHeader(http.StatusConflict)
		json.NewEncoder(w).Encode(map[string]any{
			"Message": "Produto já existe no carrinho",
		})
		return
	}

	_, err = c.cartRepository.InsertCartProductByCartId(int64(id), input)
	if err != nil {
		log.Printf("Error ao adicionar produto no carrinho: %v", err)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]any{
		"Message": "Produto adicionado carrinho com sucesso!",
	})
}
