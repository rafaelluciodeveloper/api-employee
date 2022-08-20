package main

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"go-api/configs"
	"go-api/handlers"
	"net/http"
	"os"
)

func main() {
	err := configs.Load()
	if err != nil {
		panic(err)
	}

	r := chi.NewRouter()
	r.Post("/", handlers.Create)
	r.Put("/{id}", handlers.Update)
	r.Delete("/{id}", handlers.Update)
	r.Get("/", handlers.List)
	r.Get("/{id}", handlers.Get)

	http.ListenAndServe(fmt.Sprintf(":%s", os.Getenv("PORT")), r)
}
