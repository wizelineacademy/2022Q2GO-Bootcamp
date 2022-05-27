package controller

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/luischitala/2022Q2GO-Bootcamp/models"
	"github.com/luischitala/2022Q2GO-Bootcamp/repository"
	"github.com/luischitala/2022Q2GO-Bootcamp/usecase"
)

type CharResponse struct {
	Message string `json:"message"`
	Status  bool   `json:"status"`
}

type cc struct {
	cu usecase.Character
}

type CharacterController interface {
	ListCharacter(response http.ResponseWriter, request *http.Request)
	ListCharacterApi(response http.ResponseWriter, request *http.Request)
	ReadCsv(response http.ResponseWriter, request *http.Request)
}

func NewCharacterController(cu usecase.Character) CharacterController {
	return &cc{
		cu,
	}
}

func (c *cc) ListCharacter(w http.ResponseWriter, r *http.Request) {
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

func (c *cc) ListCharacterApi(w http.ResponseWriter, r *http.Request) {
	// Client
	cli := http.Client{Timeout: time.Duration(1) * time.Second}
	req, err := http.NewRequest("GET", "https://rickandmortyapi.com/api/character", nil)
	if err != nil {
		fmt.Printf("error %s", err)
		return
	}
	req.Header.Add("Accept", `application/json`)
	// Request exec
	resp, err := cli.Do(req)

	if err != nil {
		log.Fatalln(err)
	}
	// Close response
	defer resp.Body.Close()

	bodyBytes, _ := ioutil.ReadAll(resp.Body)

	// Convert response body to string
	bodyString := string(bodyBytes)
	fmt.Println("API Response as String:\n" + bodyString)

	// Convert response body to CaracterResponse struct
	var charStruct models.CharacterResponse
	json.Unmarshal(bodyBytes, &charStruct)
	fmt.Printf("API Response as struct %+v\n", charStruct)

	//Create Csv
	csvFile, err := os.Create("files/characterResult.csv")

	if err != nil {
		log.Fatalf("failed creating file: %s", err)
	}

	// Write into the csv
	csvwriter := csv.NewWriter(csvFile)

	// Close and flush the csv
	defer csvFile.Close()
	defer csvwriter.Flush()

	// Iterate the response
	for _, character := range charStruct.Results {
		// Populate the row
		row := []string{
			strconv.Itoa(character.Id),
			character.Name,
			character.Status,
			character.Species,
			character.Type,
			character.Gender,
			character.Image,
			character.Url,
			character.Created,
		}
		// Write the row
		if err := csvwriter.Write(row); err != nil {
			log.Fatalln("error writing record to file", err)
		}
	}
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(charStruct)
}

func (c *cc) ReadCsv(w http.ResponseWriter, r *http.Request) {
	characters, err := c.cu.ReadCsv()
	if err != nil {
		fmt.Println(err)
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(characters)
}
