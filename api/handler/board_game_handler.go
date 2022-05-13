package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"

	"github.com/esvarez/go-api/internal/boardgame"
	errs "github.com/esvarez/go-api/pkg/error"
)

const boardGameID = "board_game_id"

type BoardGameHandler struct {
	BoardGameService boardgame.UseCase
}

func NewBoardGameHandler(service boardgame.UseCase) *BoardGameHandler {
	return &BoardGameHandler{
		BoardGameService: service,
	}
}

func (b *BoardGameHandler) findBoardGameByID() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		id, err := strconv.Atoi(params[boardGameID])
		if err != nil {
			log.Printf("error gettin board game id: %v", err)
			// TODO refactor response
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Internal server error"))
			return
		}
		bg, err := b.BoardGameService.FindByID(id)
		if err != nil {
			log.Printf("error getting boardgame: %v", err)
			switch {
			case errors.Is(err, errs.ErrNotFound):
				w.WriteHeader(http.StatusNotFound)
				w.Write([]byte(fmt.Sprintf("BoardGame with id %d not found", id)))
			default:
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte("Internal server error"))
			}
			return
		}

		response, err := json.Marshal(bg)
		if err != nil {
			log.Printf("error marshalling boardgame: %v", err)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Internal server error"))
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write(response)
	})
}

func MakeBoardGameHandler(router *mux.Router, handler *BoardGameHandler) {
	router.Handle("/boardgame/{board_game_id}", handler.findBoardGameByID()).
		Methods("GET")
}
