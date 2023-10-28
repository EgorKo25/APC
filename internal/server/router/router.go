package router

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type Handlers interface {
	MainPage(w http.ResponseWriter, _ *http.Request)
	GetTasksList(w http.ResponseWriter, _ *http.Request)
	SetTaskToQueue(w http.ResponseWriter, r *http.Request)
}

func NewRouter(handlers Handlers) chi.Router {

	mux := chi.NewRouter()

	mux.Use(middleware.Logger)

	mux.Route("/", func(r chi.Router) {
		r.Get("/get", handlers.GetTasksList)
		r.Get("/", handlers.MainPage)

		r.Post("/set", handlers.SetTaskToQueue)
	})

	return mux
}
