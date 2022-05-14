package main

import (
	"encoding/csv"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
)

var pokemons = map[int64]string{}

func UnmarshalData(rows [][]string) {
	for _, row := range rows {
		id, _ := strconv.ParseInt(row[0], 0, 0)
		pokemons[id] = row[1]

	}

}

func ReadCsv(filename string) ([][]string, error) {

	//Open CSV
	file, err := os.Open(filename)
	if err != nil {
		log.Println("Cannot open CSV file:", err)
		return [][]string{}, err
	}

	defer file.Close()

	reader := csv.NewReader(file)
	rows, err := reader.ReadAll()

	if err != nil {
		log.Println("Cannot read CSV file:", err)
	}

	return rows, nil
}

func handlerGetPokemonID(w http.ResponseWriter, r *http.Request) {
	res, _ := ReadCsv("data.csv")
	UnmarshalData(res)
	log.Println(pokemons)
	log.Println("The URL that you are calling is: " + r.URL.Path)
	ids := strings.TrimPrefix(r.URL.Path, "/pokemon/")

	id, _ := strconv.ParseInt(ids, 0, 0)
	log.Printf("id %T", id)
	val, ok := pokemons[id]
	if !ok {
		log.Println("id is invalid")
		w.Write([]byte("id is invalid"))
	}
	w.Write([]byte(val))
}

func main() {

	m := http.NewServeMux()
	m.HandleFunc("/pokemon/", handlerGetPokemonID)

	s := &http.Server{
		Addr:    ":8000",
		Handler: m,
	}

	log.Fatal(s.ListenAndServe())

}
