package main

import (
	"log"
	"net/http"
	"os"
	"wizeline/ghibli/adapters/controllers"
	"wizeline/ghibli/adapters/csv"
)

func main() {
	data, _ := os.Open("data/ghibli.csv")
	router := http.NewServeMux()
	filmCotroller := controllers.NewFilmController(
		csv.NewCsvRepository(
			data,
			csv.NewFilmStructurer(),
		),
	)

	router.HandleFunc("/films", filmCotroller.ServeFilms)

	log.Fatal(http.ListenAndServe(":8009", router))
}
