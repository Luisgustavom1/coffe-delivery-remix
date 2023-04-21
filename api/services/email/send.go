package email

import (
	"coffee-delivery-remix/api/models"
	"fmt"
	"net/smtp"
	"os"
)

func SendMail(email models.Email) error {
	from := os.Getenv("EMAIL_SENDER")
	password := os.Getenv("EMAIL_PASSWORD")

	smtpHost := os.Getenv("EMAIL_HOST")
	smtpPort := os.Getenv("EMAIL_PORT")

	auth := LoginAuth(from, password)

	message := generateEmailMessage(email, from)

	err := smtp.SendMail(
		smtpHost+":"+smtpPort,
		auth,
		from,
		[]string{email.To},
		message,
	)

	return err
}

func generateEmailMessage(e models.Email, from string) []byte {
	return []byte(fmt.Sprintf("From: %s\nTo: %s\nSubject: %s\n\n\n%s", from, e.To, e.Subject, e.Message))
}
