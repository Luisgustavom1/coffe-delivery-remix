package cart

import (
	"coffe-delivery-remix/api/entities"
	cart "coffe-delivery-remix/api/pkg/controller/cart/models"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

func Get(w http.ResponseWriter, request *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(request, "id"))
	if err != nil {
		log.Printf("Erro ao fazer o parse do query param id: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	cart, err := cart.GetBy(int64(id))
	var response *entities.ProductCartSimple

	if err != nil && cart.ID != 0 {
		log.Printf("Erro ao trazer registro: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	if cart.ID != 0 {
		response = &cart
	} else {
		response = (*entities.ProductCartSimple)(nil)
	}

	w.Header().Add("Content-type", "application/json")
	json.NewEncoder(w).Encode(response)
}
