package main

import (
	"encoding/csv"
	"fmt"
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
	ids := strings.TrimPrefix(r.URL.Path, "/pokemon/")

	if len(ids) == 0 {
		fmt.Println("id is missing in parameters")
		w.Write([]byte("id is missing in parameters"))
	}

	id, _ := strconv.ParseInt(ids, 0, 0)
	val, ok := pokemons[id]
	if !ok {
		log.Println("id is invalid")
		w.Write([]byte("id is invalid"))
	}

	w.Write([]byte(val + "\n"))
}

func main() {

	m := http.NewServeMux()
	m.HandleFunc("/pokemon/", handlerGetPokemonID)

	s := &http.Server{
		Addr:    ":8080",
		Handler: m,
	}

	log.Fatal(s.ListenAndServe())

}
