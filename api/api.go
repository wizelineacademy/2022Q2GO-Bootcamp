package api

import (
	"log"

	"github.com/esvarez/go-api/api/handler"
	"github.com/esvarez/go-api/infrastructure/csv"
	"github.com/esvarez/go-api/internal/boardgame"
	"github.com/esvarez/go-api/internal/pokemon"

	"github.com/gorilla/mux"
)

// TODO move to config file
const (
	Port          = "4200"
	bgPathDB      = "board_games_db.csv"
	pokemonPathDB = "pokemon_db.csv"
)

func Start() {
	var (
		router     = mux.NewRouter()
		server     = newServer(Port, router)
		csvClient  = csv.NewCSVClient(bgPathDB)
		pokemonCSV = csv.NewCSVClient(pokemonPathDB)

		bgService   = boardgame.NewService(csvClient)
		pokeService = pokemon.NewService(pokemonCSV)

		bgHandler   = handler.NewBoardGameHandler(bgService)
		pokeHandler = handler.NewPokemonHandler(pokeService)
	)

	handler.MakeBoardGameHandler(router, bgHandler)
	handler.MakePokemonHandler(router, pokeHandler)

	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("%v error starting server", err)
	}
}
