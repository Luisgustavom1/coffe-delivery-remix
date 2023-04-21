package checkout

import (
	"coffee-delivery-remix/api/entities"
	cart "coffee-delivery-remix/api/pkg/controller/cart/models"
	"coffee-delivery-remix/api/pkg/controller/errors"
	"coffee-delivery-remix/api/pkg/serialize"
	"coffee-delivery-remix/api/services/email"
	"strconv"

	"encoding/json"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func Checkout(w http.ResponseWriter, request *http.Request) {
	cartId, err := strconv.Atoi(chi.URLParam(request, "id"))
	if err != nil {
		log.Printf("Erro ao fazer o parse do query param id: %v", err)
		http_error.HttpError(w, 500)
		return
	}

	_, err = cart.DeleteBy(int64(cartId))
	if err != nil {
		log.Printf("Erro ao fazer checkout: %v", err)
		http_error.HttpError(w, 500)
		return
	}

	var modelEmail entities.EmailSimple
	err = json.NewDecoder(request.Body).Decode(&modelEmail)
	if err != nil {
		log.Printf("Erro ao decodificar o json: %v", err)
		http_error.HttpError(w, 500)
		return
	}

	var emailSerialized entities.Email
	serialize.Email(modelEmail, &emailSerialized)

	err = email.SendMail(emailSerialized)
	if err != nil {
		log.Printf("Erro ao enviar o email: %v", err)
		http_error.HttpError(w, 500)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]any{
		"Message": "Checkout feito com sucesso! Verifique seu e-mail",
	})
}
