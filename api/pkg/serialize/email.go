package serialize

import (
	"coffee-delivery-remix/api/models"
)

func Email(email models.EmailSimple, emailSerialized *models.Email) {
	*emailSerialized = models.Email{
		Message: []byte(email.Message),
		To:      email.To,
		Subject: email.Subject,
	}
}
