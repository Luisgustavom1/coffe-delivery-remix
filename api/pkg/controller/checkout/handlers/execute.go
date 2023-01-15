package checkout

import (
	"coffe-delivery-remix/api/pkg/controller/cart/models"

	"coffe-delivery-remix/api/models"
	"coffe-delivery-remix/api/pkg/serialize"
	"coffe-delivery-remix/api/services/email"

	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

func Checkout(w http.ResponseWriter, request *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(request, "id"))
	if err != nil {
		log.Printf("Erro ao fazer o parse do query param id: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	rowsAffected, err := cart.DeleteBy(int64(id))
	if err != nil {
		log.Printf("Erro ao deletar carrinho: %v", id)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	if rowsAffected == 0 {
		log.Printf("Carrinho não encontrado")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	var modelEmail models.EmailSimple
	err = json.NewDecoder(request.Body).Decode(&modelEmail)

	if err != nil {
		log.Printf("Erro ao decodificar o json: %v", err)
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
