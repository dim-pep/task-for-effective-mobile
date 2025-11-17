package web

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func CreateRouter() *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Post("/subscriptions", postSub)

	r.Get("/subscriptions/{id}", getSub)

	r.Put("/subscriptions/{id}", putSub)

	r.Delete("/subscriptions/{id}", delSub)

	r.Get("/subscriptions/list", getListSub) // List

	r.Get("/subscriptions/sum", getSumFiltredListOfSub)
	return r
}
