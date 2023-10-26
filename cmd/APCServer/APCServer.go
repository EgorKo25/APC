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

	sched, err := scheduler.NewScheduler(int32(conf.StorageInterval), int32(conf.QMax), conf.StoragePath)
	if err != nil {
		log.Fatalf("scheduler create error: \"%s\"", err)
	}

	go sched.Run()

	handlers := handler.NewHandler(sched)

	mux := router.NewRouter(handlers)

	log.Println(http.ListenAndServe(conf.ServerAddr, mux))
}
