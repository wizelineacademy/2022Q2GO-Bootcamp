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
		return nil, fmt.Errorf("can not open the file, err is %+v", err)
	}

	defer fs.Close()

	reader := csv.NewReader(fs)

	for {
		content, err := reader.Read()

		if err != nil && err != io.EOF {
			return nil, fmt.Errorf("can not read, err is %+v", err)
		}
		if err == io.EOF {
			break
		}

		if content[0] == pokemonId {
			return content, nil
		}
	}

	return nil, fmt.Errorf("Not found, err is %+v", err)
}

func handler(response http.ResponseWriter, request *http.Request) {
	keys, ok := request.URL.Query()["id"]

	if !ok || len(keys[0]) < 1 {
		log.Fatal("Url Param 'id' is missing")
		return
	}

	key := keys[0]

	content, errCsv := readCsv(key)

	if errCsv != nil {
		log.Fatal(fmt.Errorf("can not read the csv file: %v", errCsv))
	}

	jsonEncoder := json.NewEncoder(response)
	jsonEncoder.SetIndent("", "  ")

	err := jsonEncoder.Encode(content)

	if err != nil {
		err = fmt.Errorf("impossible to encode pokemons: %v", err)
		http.Error(response, err.Error(), http.StatusInternalServerError)
		log.Print(err)
	}
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe("localhost:8080", nil)
}
