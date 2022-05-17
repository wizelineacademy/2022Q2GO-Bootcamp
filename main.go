package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

const fileName = "pokemon.csv"

func readCsv(pokemonId string) ([]string, error) {
	fs, err := os.Open(fileName)

	if err != nil {
		return nil, fmt.Errorf("Can not open the file, err is %+v", err)
	}

	defer fs.Close()

	reader := csv.NewReader(fs)

readLoop:
	for {
		content, err := reader.Read()

		switch {
		case err != nil && err != io.EOF:
			return nil, fmt.Errorf("Can not read, err is %+v", err)
		case err == io.EOF:
			break readLoop
		case content[0] == pokemonId:
			return content, nil
		}
	}

	return nil, fmt.Errorf("Pokemon not found")
}

func handler(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")

	keys, ok := request.URL.Query()["id"]

	if !ok || len(keys[0]) < 1 {
		errParameterId := fmt.Errorf("Url Param 'id' is missing")
		http.Error(response, errParameterId.Error(), http.StatusBadRequest)
		log.Print(errParameterId)
		return
	}

	key := keys[0]

	content, errCsv := readCsv(key)

	if errCsv != nil {
		http.Error(response, errCsv.Error(), http.StatusNotFound)
		log.Print(errCsv)
		return
	}

	jsonEncoder := json.NewEncoder(response)
	jsonEncoder.SetIndent("", "  ")

	err := jsonEncoder.Encode(content)

	if err != nil {
		err = fmt.Errorf("Impossible to encode pokemons: %v", err)
		http.Error(response, err.Error(), http.StatusInternalServerError)
		log.Print(err)
	}
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe("localhost:8080", nil)
}
