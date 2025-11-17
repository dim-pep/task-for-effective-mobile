package web

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	httpSwagger "github.com/swaggo/http-swagger"
)

func CreateRouter() *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Get("/swagger/*", httpSwagger.WrapHandler)

	r.Post("/subscriptions", postSub)

	r.Get("/subscriptions/{id}", getSub)

	r.Put("/subscriptions/{id}", putSub)

	r.Delete("/subscriptions/{id}", delSub)

	r.Get("/subscriptions/list", getListSub) // List

	r.Post("/subscriptions/sum", getSumFiltredListOfSub)
	return r
}
