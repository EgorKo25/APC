package router

import (
	"github.com/EgorKo25/APC/internal/server/handler"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func NewRouter(handlers *handler.Handler) chi.Router {

	mux := chi.NewRouter()

	mux.Use(middleware.Logger)

	mux.Route("/", func(r chi.Router) {
		r.Get("/get", handlers.GetTasksList)
		r.Get("/", handlers.MainPage)

		r.Post("/set", handlers.SetTaskToQueue)
	})

	return mux
}
