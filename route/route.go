package route

import (
	"github.com/go-chi/chi/v5"
	httpSwagger "github.com/swaggo/http-swagger"
	"go-api/handlers"
)

func SetupRoute() *chi.Mux {
	r := chi.NewRouter()
	r.Route("/employee", func(r chi.Router) {
		r.Post("/", handlers.Create)
		r.Put("/{id}", handlers.Update)
		r.Delete("/{id}", handlers.Delete)
		r.Get("/", handlers.List)
		r.Get("/{id}", handlers.Get)
	})
	
	r.Get("/swagger/*any", httpSwagger.WrapHandler)

	return r
}
