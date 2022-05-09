package handlers

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/luischitala/2022Q2GO-Bootcamp/models"
	"github.com/luischitala/2022Q2GO-Bootcamp/repository"
	"github.com/luischitala/2022Q2GO-Bootcamp/server"
)

type CharResponse struct {
	Message string `json:"message"`
	Status  bool   `json:"status"`
}

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
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(characters)
	}
}

// Handler that reads a csv file and return a json response with its content
func ReadCsvHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		characters := make([]models.Character, 0)

		csvFile, err := os.Open("files/characters.csv")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			json.NewEncoder(w).Encode(HomeResponse{
				Message: "Unable to read input file",
			})
			return
		}
		defer csvFile.Close()

		csvLines, err := csv.NewReader(csvFile).ReadAll()
		if err != nil {
			fmt.Println(err)
		}
		for _, line := range csvLines {
			id, _ := strconv.Atoi(line[0])

			character := models.Character{
				Id:   id,
				Name: line[1],
			}
			characters = append(characters, character)

		}
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(characters)

	}
}
