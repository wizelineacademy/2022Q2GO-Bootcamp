package controller
import (
	"fmt"
	"net/http"
	"encoding/json"
	"io/ioutil"
	"github.com/waltergomezda/2022Q2GO-Bootcamp/internal/service"
	"github.com/waltergomezda/2022Q2GO-Bootcamp/internal/entity"
	"github.com/waltergomezda/2022Q2GO-Bootcamp/internal/repository"
)
// GetHandler handles the index route
func GetGreetingHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome!, 2022Q2GO-Bootcamp"))
}

// GetCSVHandler handles the CSV route
func GetCSVHandler(w http.ResponseWriter, r *http.Request) {
	repo:= repository.NewCsvRepository("data.csv")
	svc := service.NewCsvService(repo)
	results, err := svc.ExtractCsvRecords()
	if err != nil {
		http.Error(w, "Error retrieving data",
			http.StatusInternalServerError)
	}
	jsonBody, err := json.Marshal(results)
	if err != nil {
		http.Error(w, "Error converting results to json",
			http.StatusInternalServerError)
	}
	w.Write(jsonBody)
}

// GetExternalApiHandler handles the External Api route
func GetExternalApiHandler(w http.ResponseWriter, r *http.Request) {
	response, err := http.Get("https://catfact.ninja/fact")
	if err != nil {
		http.Error(w, "Error retrieving data",
			http.StatusInternalServerError)
	}
	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil{
		http.Error(w, "Error retrieving data",
			http.StatusInternalServerError)
	}

	var catFact entity.CatFact 
	if err := json.Unmarshal(responseData, &catFact); err != nil {
        panic(err)
    }
	fmt.Println( catFact)
	repo:= repository.NewCsvRepository("externalApiData.csv")
	svc := service.NewCsvService(repo)
	err = svc.SaveCsvRecords([]entity.CatFact{catFact})
	if err != nil{
		http.Error(w, "Error retrieving data",
			http.StatusInternalServerError)
	}
	w.Write(responseData)
}
