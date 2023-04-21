package checkout

import (
	"coffee-delivery-remix/api/models"
	"coffee-delivery-remix/api/pkg/controller/cart/models"
	"coffee-delivery-remix/api/pkg/serialize"
	"coffee-delivery-remix/api/services/email"

	"encoding/json"
	"log"
	"net/http"
)

func Checkout(w http.ResponseWriter, request *http.Request) {
	_, err := cart.DeleteAll()
	if err != nil {
		log.Printf("Erro deletar cart: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	var modelEmail models.EmailSimple
	err = json.NewDecoder(request.Body).Decode(&modelEmail)
	if err != nil {
		log.Printf("Erro ao decodificar o json: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	var emailSerialized models.Email
	serialize.Email(modelEmail, &emailSerialized)

	err = email.SendMail(emailSerialized)
	if err != nil {
		log.Printf("Erro ao enviar o email: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]any{
		"Message": "Checkout feito com sucesso! Verifique seu e-mail",
	})
}
