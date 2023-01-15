package email

import (
	"coffe-delivery-remix/api/models"
	"errors"
	"net/smtp"
	"os"
)

type loginAuth struct {
	username string
	password string
}

func (a *loginAuth) Start(server *smtp.ServerInfo) (string, []byte, error) {
	return "LOGIN", []byte{}, nil
}

func (a *loginAuth) Next(fromServer []byte, more bool) ([]byte, error) {
	if more {
		switch string(fromServer) {
		case "Username:":
			return []byte(a.username), nil
		case "Password:":
			return []byte(a.password), nil
		default:
			return nil, errors.New("unknown fromServer")
		}
	}
	return nil, nil
}

func SendMail(email models.Email) error {
	from := os.Getenv("EMAIL_SENDER")
	password := os.Getenv("EMAIL_PASSWORD")

	smtpHost := os.Getenv("EMAIL_HOST")
	smtpPort := os.Getenv("EMAIL_PORT")

	auth := LoginAuth(from, password)

	err := smtp.SendMail(
		smtpHost+":"+smtpPort, 
		auth, 
		from, 
		email.To, 
		email.Message,
	)

	return err
}

func LoginAuth(username, password string) smtp.Auth {
	return &loginAuth{username, password}
}
