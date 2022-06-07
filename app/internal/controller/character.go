package controller

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/luischitala/2022Q2GO-Bootcamp/internal/repository"
	"github.com/luischitala/2022Q2GO-Bootcamp/internal/usecase"
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
	ReadCsvConcurrently(response http.ResponseWriter, request *http.Request)
}

func NewCharacterController(cu usecase.Character) CharacterController {
	return &cc{
		cu,
	}
}

// ListCharacter godoc
// @Summary Character List DB
// @Description Read Characters from database
// @Tags Character
// @Accept  json
// @Produce  json
// @Success 200
// @Router /characters [get]
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

// ListCharacter godoc
// @Summary Character List API
// @Description Read Characters from the Rick & Morty's API
// @Tags Character
// @Accept  json
// @Produce  json
// @Success 200
// @Router /charactersApi [get]
func (c *cc) ListCharacterApi(w http.ResponseWriter, r *http.Request) {

	charStruct, err := c.cu.WriteCsv()
	if err != nil {
		log.Fatalln(err)
	}
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(charStruct)
}

// ReadCsvCharacter godoc
// @Summary Character List Csv
// @Description Read Characters from a Csv File
// @Tags Character
// @Accept  json
// @Produce  json
// @Success 200
// @Router /readCsv [get]
func (c *cc) ReadCsv(w http.ResponseWriter, r *http.Request) {
	characters, err := c.cu.ReadCsv()
	if err != nil {
		fmt.Println(err)
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(characters)
}

// ReadCsvCharacterConcurrently godoc
// @Summary Character List Csv
// @Description Read Characters from a Csv File Concurrently
// @Param type query string true  "odd or even"
// @Param items query string true  "Quantity of items to retrieve"
// @Param items_per_worker query string true  "Quantity of items that each worker will read concurrently"
// @Tags Character
// @Accept  json
// @Produce  json
// @Success 200
// @Router /readCsvConcurrently [get]
func (c *cc) ReadCsvConcurrently(w http.ResponseWriter, r *http.Request) {

	typeP := r.URL.Query().Get("type")
	// Validate that the type parameter has been sent
	if len(typeP) == 0 {
		json.NewEncoder(w).Encode("Must indicate if the result will be odd or even")
		return
	}
	itemsP := r.URL.Query().Get("items")
	// Validate that the items parameter has been sent
	if len(itemsP) == 0 {
		json.NewEncoder(w).Encode("Must indicate the quantity of items to retrieve")
		return
	}

	// Validate that the items parameter is a numeric value
	items, err := strconv.Atoi(itemsP)
	if err != nil {
		json.NewEncoder(w).Encode("Please enter a numeric value for the 'items' parameter")
		return
	}

	itemsPerWorkerP := r.URL.Query().Get("items_per_worker")
	// Validate that the items_per_worker parameter has been sent
	if len(itemsPerWorkerP) == 0 {
		json.NewEncoder(w).Encode("Must send the quantity of items for the workers to retrieve")
		return
	}
	// Validate that the items parameter is a numeric value
	itemsPerWorker, err := strconv.Atoi(itemsPerWorkerP)
	if err != nil {
		json.NewEncoder(w).Encode("Please enter a numeric value for the 'items_per_worker' parameter")
		return
	}
	//Call the Csv Reader Concurrently using the query parameters to adjust the workload
	characters, err := c.cu.ReadCsvConcurrently(typeP, items, itemsPerWorker)
	if err != nil {
		fmt.Println(err)
	}
	// Parse the response
	response, err := json.Marshal(struct {
		Code int         `json:"code"`
		Data interface{} `json:"data"`
	}{http.StatusOK, characters})

	if err != nil {
		json.NewEncoder(w).Encode("Error while reading the CSV file")
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)

}
