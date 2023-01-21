package main

import (
	"coffe-delivery-remix/api/cmd/routes"
	"coffe-delivery-remix/api/configs"

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

	c := cors.New(cors.Options{
		AllowCredentials: true,
		AllowedMethods:   []string{"POST", "PUT", "DELETE", "GET"},
	})

	r := chi.NewRouter()

	r.Use(c.Handler)

	r.Mount("/v1", routes.Routes())

	fmt.Printf("Starting server in the :%s\n\n", configs.GetApiServerPort())
	http.ListenAndServe(fmt.Sprintf(":%s", configs.GetApiServerPort()), r)
}
