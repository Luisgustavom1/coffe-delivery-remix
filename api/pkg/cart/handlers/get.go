package cart

import (
	"coffe-delivery-remix/api/entities"
	cart "coffe-delivery-remix/api/pkg/cart/models"
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
	var response *entities.CartProductSimple

	if err != nil && cart.ID != 0 {
		log.Printf("Erro ao trazer registro: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	if cart.ID != 0 {
		response = &cart
	} else {
		response = (*entities.CartProductSimple)(nil)
	}

	w.Header().Add("Content-type", "application/json")
	json.NewEncoder(w).Encode(response)
}
