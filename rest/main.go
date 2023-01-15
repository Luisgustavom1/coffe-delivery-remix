package main

import (
	"coffe-delivery-remix/rest/cmd/routes"
	"coffe-delivery-remix/rest/configs"
	"coffe-delivery-remix/rest/api/middleware"

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

	middleware.Cors(r)

	r.Mount("/v1", routes.Routes())

	http.ListenAndServe(fmt.Sprintf(":%s", configs.GetApiServerPort()), r)
}
