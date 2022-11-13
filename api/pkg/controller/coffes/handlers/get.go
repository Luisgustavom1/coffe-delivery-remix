package coffes

import (
	"coffe-delivery-remix/api/entities"
	coffes "coffe-delivery-remix/api/pkg/controller/coffes/models"
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

	coffe, err := coffes.GetBy(int64(id))
	var response *entities.Coffe[string]

	if err != nil && coffe.ID != 0 {
		log.Printf("Erro ao trazer registro: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	if coffe.ID != 0 {
		response = &coffe
	} else {
		response = (*entities.Coffe[string])(nil)
	}

	w.Header().Add("Content-type", "application/json")
	json.NewEncoder(w).Encode(response)
}
