package controller

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	models "github.com/luischitala/2022Q2GO-Bootcamp/internal/entity"
	"github.com/luischitala/2022Q2GO-Bootcamp/internal/repository"
	"github.com/luischitala/2022Q2GO-Bootcamp/internal/usecase"
)

type CharResponse struct {
	Code int                 `json:"code"`
	Data []*models.Character `json:"data"`
}

type cc struct {
	cu usecase.Character
}

type CharacterController interface {
	ListCharacter(response http.ResponseWriter, request *http.Request)
	GetCharactersAndWriteOnCsv(response http.ResponseWriter, request *http.Request)
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
// @Param id query string false "Retrieve a Character by id"
// @Accept  json
// @Produce  json
// @Success 200
// @Router /characters [get]
func (c *cc) ListCharacter(w http.ResponseWriter, r *http.Request) {
	var err error
	id := r.URL.Query().Get("id")
	var parsedId = uint64(0)
	if id != "" {
		parsedId, err = strconv.ParseUint(id, 10, 53)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

	} else {
		parsedId = 0
	}
	characters, err := repository.ListCharacter(r.Context(), parsedId)
	log.Println(characters)

	if err != nil {
		HandleError(w, http.StatusInternalServerError, err.Error())
		return
	}
	// Parse the response
	response := CharResponse{
		Code: http.StatusOK,
		Data: characters,
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

// ListCharacter godoc
// @Summary Character List API
// @Description Read Characters from the Rick & Morty's API and wirtes the response into a CSV File
// @Tags Character
// @Accept  json
// @Produce  json
// @Success 200
// @Router /charactersApi [get]
func (c *cc) GetCharactersAndWriteOnCsv(w http.ResponseWriter, r *http.Request) {
	// Call the external API using ConsultApi method that returns a characters response in []byte format
	apiResponse, err := c.cu.ConsultApi()
	if err != nil {
		log.Fatalln(err)
	}
	// Write the result of the api request, on the csv if it exists
	charStruct, err, _ := c.cu.WriteCsv(apiResponse)
	if err != nil {
		log.Println(err)
		HandleError(w, http.StatusInternalServerError, "Error while writing CSV file")
		return
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
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
		log.Println(err)
		HandleError(w, http.StatusInternalServerError, "Error while reading CSV file")
		return
	}
	// Parse the response

	response := CharResponse{
		Code: http.StatusOK,
		Data: characters,
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

// ReadCsvCharacterConcurrently godoc
// @Summary Character List Csv
// @Description Read Characters from a Csv File Concurrently
// @Param type query string true   "Retrieve 'odd' or 'even' ids"
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
		log.Println("No type parameter sent")
		HandleError(w, http.StatusBadRequest, "Must indicate if the result will be odd or even")
		return
	}
	itemsP := r.URL.Query().Get("items")
	// Validate that the items parameter has been sent
	if len(itemsP) == 0 {
		log.Println("No items parameter sent")
		HandleError(w, http.StatusBadRequest, "Must indicate the quantity of items to retrieve")
		return
	}

	// Validate that the items parameter is a numeric value
	items, err := strconv.Atoi(itemsP)
	if err != nil {
		log.Println("Wrong type for items parameter")
		HandleError(w, http.StatusBadRequest, "Please enter a numeric value for the 'items' parameter")
		return
	}

	itemsPerWorkerP := r.URL.Query().Get("items_per_worker")
	// Validate that the items_per_worker parameter has been sent
	if len(itemsPerWorkerP) == 0 {
		log.Println("No items_per_worker parameter sent")
		HandleError(w, http.StatusBadRequest, "Must send the quantity of items for the workers to retrieve")
		return
	}
	// Validate that the items parameter is a numeric value
	itemsPerWorker, err := strconv.Atoi(itemsPerWorkerP)
	if err != nil {
		log.Println("Wrong type for items_per_worker parameter")
		HandleError(w, http.StatusBadRequest, "Please enter a numeric value for the 'items_per_worker' parameter")
		return
	}
	//Call the Csv Reader Concurrently using the query parameters to adjust the workload
	characters, err := c.cu.ReadCsvConcurrently(typeP, items, itemsPerWorker)
	if err != nil {
		log.Println(err)
		HandleError(w, http.StatusInternalServerError, err.Error())
		return

	}
	// Parse the response
	response, err := json.Marshal(struct {
		Code int         `json:"code"`
		Data interface{} `json:"data"`
	}{http.StatusOK, characters})

	if err != nil {
		log.Println(err)
		HandleError(w, http.StatusInternalServerError, err.Error())
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)

}

func HandleError(w http.ResponseWriter, code int, message string) {
	response, _ := json.Marshal(struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
	}{code, message})
	w.WriteHeader(code)
	w.Write(response)

}
