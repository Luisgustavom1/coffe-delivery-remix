package middleware

import (
	"github.com/go-chi/chi/v5"
	"github.com/rs/cors"
)

func Cors(router *chi.Mux) {
	c := cors.New(cors.Options{
		AllowCredentials: true,
		AllowedMethods:   []string{"POST", "PUT", "DELETE", "GET"},
	})

	router.Use(c.Handler)
}