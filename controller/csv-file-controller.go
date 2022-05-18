package controller

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/SamuelSalas/2022Q2GO-Bootcamp/errors"
	"github.com/SamuelSalas/2022Q2GO-Bootcamp/service"
)

type CSVFileController interface {
	PostCSVFile(resp http.ResponseWriter, req *http.Request)
}

type controller struct{}

var csvService service.CsvService

func NewCsvController(service service.CsvService) CSVFileController {
	csvService = service
	return &controller{}
}

func (*controller) PostCSVFile(resp http.ResponseWriter, req *http.Request) {
	resp.Header().Set("Content-type", "application/json")
	file, fileHeader, fileError := req.FormFile("csv")

	if fileError != nil {
		resp.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(resp).Encode(errors.ErrorMessage{Message: fileError.Error()})
		return
	}

	if fileHeader.Header.Get("Content-Type") != "text/csv" {
		resp.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(resp).Encode(errors.ErrorMessage{Message: "Invalid file type"})
		return
	}

	csvReader := csv.NewReader(file)
	data, csvFile := csvReader.ReadAll()
	if csvFile != nil {
		resp.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(resp).Encode(errors.ErrorMessage{Message: csvFile.Error()})
		return
	}

	result, err := csvService.ConvertCsvToJson(data)
	fmt.Println(err.Error())
	if err != nil {
		resp.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(resp).Encode(errors.ErrorMessage{Message: err.Error()})
		return
	}

	resp.Header().Set("Content-type", "application/json")
	resp.WriteHeader(http.StatusOK)
	json.NewEncoder(resp).Encode(result)
}
