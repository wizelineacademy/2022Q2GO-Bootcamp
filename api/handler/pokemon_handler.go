package handler

import (
	"github.com/esvarez/go-api/internal/pokemon"
	"github.com/esvarez/go-api/pkg/web"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

const (
	pokemonID = "pokemon_id"
)

type PokemonHandler struct {
	pokemonService pokemonService
}

type pokemonService interface {
	FindByID(id string) (*pokemon.Pokemon, error)
	GetPokemon(tpe string, items, itemsWorker int) ([]pokemon.Pokemon, error)
}

func NewPokemonHandler(service pokemonService) *PokemonHandler {
	return &PokemonHandler{
		pokemonService: service,
	}
}

func (p PokemonHandler) findPokemon() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)

		poke, err := p.pokemonService.FindByID(params[pokemonID])
		if err != nil {
			log.Printf("error getting pokemon: %v", err)
			web.ErrorResponse(err).Send(w)
			return
		}

		web.Success(poke, http.StatusOK).Send(w)
	})
}

func (p PokemonHandler) getPokemon() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		params := r.URL.Query()
		t, webErr := getValidType(params)
		if webErr != nil {
			log.Println("invalid type")
			webErr.Send(w)
			return
		}

		itms, webErr := getValidItems(params)
		if webErr != nil {
			webErr.Send(w)
			return
		}

		itmsWorker, webErr := getValidItemsWorkers(params)
		if webErr != nil {
			webErr.Send(w)
			return
		}

		response, err := p.pokemonService.GetPokemon(t, itms, itmsWorker)
		if err != nil {
			web.InternalServerError.Send(w)
			return
		}

		web.Success(response, http.StatusOK).Send(w)
	})
}

func MakePokemonHandler(r *mux.Router, handler *PokemonHandler) {
	r.Handle("/pokemon/{pokemon_id}", handler.findPokemon()).
		Methods(http.MethodGet)
	r.Handle("/pokemon", handler.getPokemon()).
		Methods(http.MethodGet)
}
