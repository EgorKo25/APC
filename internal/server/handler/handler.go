package handler

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/EgorKo25/APC/internal/scheduler"
)

// NewHandler constructor Handler
func NewHandler(s *scheduler.Scheduler) *Handler {
	return &Handler{s: s}
}

type Handler struct {
	s *scheduler.Scheduler
}

// MainPage return / page
func (h *Handler) MainPage(w http.ResponseWriter, _ *http.Request) {

	helloMsg := `<h1>Welcome to Arithmetic Progressive Calculator</h1>`

	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte(helloMsg))
}

// SetTaskToQueue set list of tasks to queue Scheduler
func (h *Handler) SetTaskToQueue(w http.ResponseWriter, r *http.Request) {

	var tasks []scheduler.Task

	body, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer func() { _ = r.Body.Close() }()

	err = json.Unmarshal(body, &tasks)
	if err != nil {
		log.Println(body)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	for _, task := range tasks {
		h.s.InsertTask(&task)
	}

	log.Println(h.s.GetSortQueue()[0])
	w.WriteHeader(http.StatusOK)
}

// GetTasksList return all of tasks
func (h *Handler) GetTasksList(w http.ResponseWriter, _ *http.Request) {

	w.Header().Add("Content-Type", "application/json")

	tasks := h.s.GetSortQueue()

	b, err := json.Marshal(tasks)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	_, err = w.Write(b)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	return
}
