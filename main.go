package main

import (
	"log"
	"net/http"

	"github.com/dalebradley/expenser-api/conf"
	"github.com/dalebradley/expenser-api/routes"
	"github.com/davecgh/go-spew/spew"
	"github.com/gorilla/mux"
)

func main() {
	cfg := conf.Get()
	spew.Dump(cfg)
	spew.Dump("Here we go again!")

	mainRouter := mux.NewRouter()
	mainRouter.StrictSlash(true)
	routes.Register(mainRouter)

	log.Println("Starting server on port " + cfg.ExpenserPort[1:])
	err := http.ListenAndServe(cfg.ExpenserPort, mainRouter)
	if err != nil {
		log.Fatal("Error starting server", err)
	}

}
