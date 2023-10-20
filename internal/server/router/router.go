package router

import (
	"net/http"

	"github.com/EgorKo25/APC/internal/server/handler"
)

func NewRouter(handlers *handler.Handler) *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", handlers.MainPage)
	mux.HandleFunc("/set/", handlers.SetTaskToQueue)
	mux.HandleFunc("/get/", handlers.GetTasksList)

	return mux
}
