package main

import (
	"coffe-delivery-remix/api/configs"
	"coffe-delivery-remix/api/handlers"
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/rs/cors"
)

func main() {
	err := configs.LoadConfigs()
	if err != nil {
		log.Printf("Erro: %v", err)
		return
	}

	r := chi.NewRouter()

	c := cors.New(cors.Options{
		AllowCredentials: true,
	})

	r.Use(c.Handler)

	r.Post("/coffes", handlers.Create)
	r.Put("/coffes/{id}", handlers.Update)
	r.Delete("/coffes/{id}", handlers.Delete)
	r.Get("/coffes", handlers.List)
	r.Get("/coffes/{id}", handlers.Get)
	r.Get("/cart", handlers.List)

	http.ListenAndServe(fmt.Sprintf(":%s", configs.GetApiServerPort()), r)
}
