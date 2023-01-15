package models

type Email struct {
	To string `json:"to"`
	Subject string `json:"subject"`
	Message []byte `json:"message"`
}

type EmailSimple struct {
	To string `json:"to"`
	Subject string `json:"subject"`
	Message string `json:"message"`
}