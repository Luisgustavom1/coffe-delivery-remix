package checkout

import (
	"coffe-delivery-remix/api/models"
	"coffe-delivery-remix/api/pkg/serialize"
	"coffe-delivery-remix/api/services/email"

	"encoding/json"
	"log"
	"net/http"
)

func Checkout(w http.ResponseWriter, request *http.Request) {
	var modelEmail models.EmailSimple
	err := json.NewDecoder(request.Body).Decode(&modelEmail)

	if err != nil {
		log.Printf("Error ao decodificar o json: %v", err)
		return
	}

	var emailSerialized models.Email
	serialize.Email(modelEmail, &emailSerialized)

	err = email.SendMail(emailSerialized)

	if err != nil {
		log.Printf("Erro ao fazer enviar o email: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]any{
		"Message": "Checkout feito com sucesso! Verifique seu e-mail",
	})
}
