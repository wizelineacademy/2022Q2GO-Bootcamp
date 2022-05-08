package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/luischitala/2022Q2GO-Bootcamp/repository"
	"github.com/luischitala/2022Q2GO-Bootcamp/server"
)

func ListCharacterHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var err error
		pageStr := r.URL.Query().Get("page")
		var page = uint64(0)
		if pageStr != "" {
			page, err = strconv.ParseUint(pageStr, 10, 53)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
		}
		characters, err := repository.ListCharacter(r.Context(), page)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(characters)
	}
}
