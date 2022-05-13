package api

import (
	"log"

	"github.com/gorilla/mux"

	"github.com/esvarez/go-api/api/handler"
	"github.com/esvarez/go-api/infrastructure/csv"
	"github.com/esvarez/go-api/internal/boardgame"
)

// TODO move to config file
const (
	Port = "4200"
	path = "db.csv"
)

func Start() {
	var (
		router    = mux.NewRouter()
		server    = newServer(Port, router)
		csvClient = csv.NewCSVClient(path)
		bgService = boardgame.NewService(csvClient)
		bgHandler = handler.NewBoardGameHandler(bgService)
	)

	handler.MakeBoardGameHandler(router, bgHandler)

	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("%v error starting server", err)
	}
}
