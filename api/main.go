package main

import (
	"coffee-delivery-remix/api/cmd/routes"
	"coffee-delivery-remix/api/configs"
	"coffee-delivery-remix/api/infra/db"

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

	connection, err := db.OpenConnection()
	if err != nil {
		log.Fatalf("Error on db connection: %v", err)
		return
	}
	defer connection.Close()

	r := chi.NewRouter()
	r.Use(c.Handler)
	r.Mount("/v1", routes.Routes(connection))

	fmt.Printf("Starting server in the :%s\n\n", configs.GetApiServerPort())
	http.ListenAndServe(fmt.Sprintf(":%s", configs.GetApiServerPort()), r)
}
