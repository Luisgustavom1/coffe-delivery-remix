package main

import (
	"coffe-delivery-remix/api/configs"
	"coffe-delivery-remix/api/cmd/routes"

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
	})

	r := chi.NewRouter()

	r.Use(c.Handler)

	r.Mount("/v1", routes.Routes())

	http.ListenAndServe(fmt.Sprintf(":%s", configs.GetApiServerPort()), r)
}
