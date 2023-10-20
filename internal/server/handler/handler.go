package handler

import (
	"net/http"
)

func NewHandler() *Handler {
	return &Handler{}
}

type Handler struct {
}

func (h *Handler) MainPage(w http.ResponseWriter, r *http.Request) {

	hellomsg := `<h1>Добро пожаловать в калькулятор вычисляющий арифметическую программию</h1>`

	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte(hellomsg))
}
func (h *Handler) SetTaskToQueue(w http.ResponseWriter, r *http.Request) {

}

func (h *Handler) GetTasksList(w http.ResponseWriter, r *http.Request) {

}
