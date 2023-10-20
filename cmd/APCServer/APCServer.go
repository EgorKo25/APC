package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/EgorKo25/APC/internal/scheduler"
	"github.com/EgorKo25/APC/internal/server/handler"
	"github.com/EgorKo25/APC/internal/server/router"
)

var (
	maxTask int
)

func main() {

	flag.IntVar(&maxTask, "max", 6, "")
	flag.Parse()

	sched := scheduler.NewScheduler(uint(maxTask))

	go sched.Run()

	handlers := handler.NewHandler()

	mux := router.NewRouter(handlers)
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatalf("error: %s", err)
	}
}
