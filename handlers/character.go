package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/TanZng/toh-api/repository"
	"github.com/TanZng/toh-api/server"
	"github.com/gorilla/mux"
)

type InsertCharacterRequest struct {
	Name string `json:"name"`
	Age  uint64 `json:"age"`
}

type CharacterRequest struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
	Age  uint64 `json:"age"`
}

func GetCharacterByIdHandler(s server.Server) http.HandlerFunc {
	return func(writter http.ResponseWriter, request *http.Request) {
		params := mux.Vars(request)
		id, err := strconv.ParseInt(params["id"], 10, 64)
		if err != nil {
			http.Error(writter, err.Error(), http.StatusBadRequest)
			return
		}

		character, err := repository.GetCharacterById(request.Context(), id)
		if err != nil {
			http.Error(writter, err.Error(), http.StatusInternalServerError)
			return
		}
		writter.Header().Set("Content-Type", "application/json")
		json.NewEncoder(writter).Encode(character)
	}
}
