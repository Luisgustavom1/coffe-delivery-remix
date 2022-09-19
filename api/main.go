package main

import (
	"coffe-delivery-remix/api/configs"
	"coffe-delivery-remix/api/handlers"
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func main() {
	err := configs.LoadConfigs()
	if err != nil {
		log.Printf("Erro: %v", err)
		return
	}

	r := chi.NewRouter()
	r.Post("/", handlers.Create)
	r.Put("/{id}", handlers.Update)
	r.Delete("/{id}", handlers.Delete)
	r.Get("/", handlers.List)
	r.Get("/{id}", handlers.Get)

	http.ListenAndServe(fmt.Sprintf(":%s", configs.GetApiServerPort()), r)
}