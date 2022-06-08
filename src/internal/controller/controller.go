package controller
import (
	"fmt"
	"net/http"
	"encoding/json"
	"io/ioutil"
	"strconv"
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

type testTask struct {
	Name string
	Line int
	TaskProcessor func(int, string, int)
}
func (t testTask) Run(workerID int){
	t.TaskProcessor(workerID, t.Name, t.Line)
}

// GetCSVConcurrentlyHandler handles the CSV data concurrently
func GetCSVConcurrentlyHandler(w http.ResponseWriter, r *http.Request) {
	itemType := r.URL.Query().Get("type") 
	strItems := r.URL.Query().Get("items")
	strItemsPerWorkers := r.URL.Query().Get("items_per_workers")

	if itemType != "odd" && itemType != "even" {
		http.Error(w, "Invalid type param",
			http.StatusBadRequest)
	}

	maxItems, err := strconv.Atoi(strItems)
    if err != nil {
		fmt.Println(err)
		http.Error(w, "Invalid Items param",
		http.StatusBadRequest)
    }

	maxItemsPerWorker, err := strconv.Atoi(strItemsPerWorkers)
    if err != nil {
		http.Error(w, "Invalid ItemsPerWorkers param",
		http.StatusBadRequest)
    }

	fmt.Println("type =>", itemType)
	fmt.Println("items =>", maxItems)
	fmt.Println("itemsPerWorkers =>", maxItemsPerWorker)

 
	numberOfWorkers := 7
	processedItems :=0

	//***********************************************

	//Load records from CSV file
	repo:= repository.NewCsvRepository("data.csv")
	svc := service.NewCsvService(repo)
	data, err := svc.ExtractCsvRecords()
	if err != nil { 
	}

	//Create job and result channel
	numJobs := len(data)
    jobs := make(chan int, numJobs)
    results := make(chan string, numJobs)
	var done = make(chan bool)

	//generate workers
	for w := 1; w <= numberOfWorkers; w++ {
        go func (workerId int, jobs <-chan int, results chan<- string) {
			processedItemsPerWorker :=0
			for idx := range jobs {
				if processedItems>=maxItems || processedItemsPerWorker >= maxItemsPerWorker {
					// fmt.Println("max element reach at worker",workerId, processedItems, maxItems, processedItemsPerWorker, maxItemsPerWorker)
					break
				}
				record := data[idx-1]
				if (itemType == "odd" && record.Key%2==0) || (itemType == "even" && record.Key%2==1) {
					 continue
		 		}
				processedItems++
				processedItemsPerWorker++
				// fmt.Println("send worker:",workerId," processed line Number: ", idx ," with content:", record.Key,"|", record.Value)
				results <- fmt.Sprint(" worker:",workerId," processed line Number: ", idx ," with content:", record.Key,"|", record.Value)
				
			}
			done <-true
		}(w, jobs, results)
    }

	//load job channel, each element represent a position in Data
	for idx := 1; idx <= numJobs; idx++ {
        jobs <- idx
    }
    close(jobs)

	// wait until all workers get their jobs completed
	for i := 1; i <= numberOfWorkers; i++ {
        <- done
    }
	//Prepare response
	var content []string 
	for a := 1; a <= processedItems; a++ {
		content = append(content, <-results) 
    } 
 
 	jsonBody, err := json.Marshal(content)
	if err != nil {
		http.Error(w, "Error converting content to json",
			http.StatusInternalServerError)
	}
	w.Write(jsonBody)
}
