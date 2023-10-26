package main

import (
	"log"
	"net/http"

	"github.com/EgorKo25/APC/internal/config"
	"github.com/EgorKo25/APC/internal/scheduler"
	"github.com/EgorKo25/APC/internal/server/handler"
	"github.com/EgorKo25/APC/internal/server/router"
)

func main() {

	conf, err := config.NewConfig()
	if err != nil {
		log.Fatalf("configuration create error: \"%s\"", err)
	}

	sched := scheduler.NewScheduler(int32(conf.QMax))

	go sched.Run()

	handlers := handler.NewHandler(sched)

	mux := router.NewRouter(handlers)
	if err := http.ListenAndServe(conf.ServerAddr, mux); err != nil {
		log.Fatalf("server error: \"%s\"", err)
	}
}
