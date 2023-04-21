package serialize

import (
	"coffee-delivery-remix/api/entities"
)

func Email(email entities.EmailSimple, emailSerialized *entities.Email) {
	*emailSerialized = entities.Email{
		Message: []byte(email.Message),
		To:      email.To,
		Subject: email.Subject,
	}
}
