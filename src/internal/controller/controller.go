package controller
import (
	"fmt"
	"net/http"
	"encoding/json"
	"io/ioutil"
	"strconv"
	"time"
	"sync"
	"github.com/waltergomezda/2022Q2GO-Bootcamp/internal/service"
	"github.com/waltergomezda/2022Q2GO-Bootcamp/internal/entity"
	"github.com/waltergomezda/2022Q2GO-Bootcamp/internal/repository"
	"github.com/waltergomezda/2022Q2GO-Bootcamp/internal/workerpool"
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

	totalItems, err := strconv.Atoi(strItems)
    if err != nil {
		fmt.Println(err)
		http.Error(w, "Invalid Items param",
		http.StatusBadRequest)
    }

	totalItemsPerWorker, err := strconv.Atoi(strItemsPerWorkers)
    if err != nil {
		http.Error(w, "Invalid ItemsPerWorkers param",
		http.StatusBadRequest)
    }

	fmt.Println("type =>", itemType)
	fmt.Println("items =>", totalItems)
	fmt.Println("itemsPerWorkers =>", totalItemsPerWorker)

	//***********************************************

	//Reading CSV
	repo:= repository.NewCsvRepository("data.csv")
	svc := service.NewCsvService(repo)
	results, err := svc.ExtractCsvRecords()
	if err != nil {
		http.Error(w, "Error retrieving CSV data",
			http.StatusInternalServerError)
	}
	
	fmt.Println(totalItems, totalItemsPerWorker)

	// declare worker pool
	pool := workerpool.NewGoroutinePool(2)
    taskSize := totalItems
	taskCounterPerWorker := make(map[int]int)
	resultsPerWorker := make(map[int]string)
	fmt.Println("taksCounterPerWorker",taskCounterPerWorker, resultsPerWorker)
    
    //wait for jobs to finish
    wg := &sync.WaitGroup{}

	readTaskFn := func(workerID int, name string, line int){
		// if taskCounterPerWorker[workerID] > totalItemsPerWorker {
		// 	fmt.Println("max items reached WorkerID",workerID )
		// 	return
		// }

		fmt.Println("inputvalues:","workerID:",workerID, "name:", name, "line:", line)
		time.Sleep(time.Second)
		
		tempResult := results[line]
		taskCounterPerWorker[workerID] ++
		if _, ok := resultsPerWorker[workerID]; ok==false {
			resultsPerWorker[workerID] = ""
		}
		resultsPerWorker[workerID] = fmt.Sprint(resultsPerWorker[workerID], " Line: ", line ," key:", tempResult.Key , " value: ", tempResult.Value)
	
		if name != ""{
			fmt.Println( fmt.Sprint("worker", workerID ," finished ", name))
		}
		wg.Done()
    }
    var tasks []testTask
	validTaskSize :=0
	
	//generate tasks considering filter type
    for v:=0; v < taskSize; v++{
		if v > len(results){
			break
		}
		if (itemType == "odd" && v%2==0) || (itemType == "even" && v%2==1) {
			continue
		}
        tasks = append(tasks, testTask{
            Name: fmt.Sprint("task ", v ),
			Line: v,
            TaskProcessor: readTaskFn,
        })
		validTaskSize++
    }
	wg.Add(validTaskSize)
    for _, task := range tasks{
        pool.ScheduleWork(task)
    }
    pool.Close()
    wg.Wait()

	//Prepare response
	var content []string 
	for workerId, value := range resultsPerWorker {
        fmt.Println("Worker",workerId, ":", value)
		content = append(content,fmt.Sprint("Worker",workerId, ":", value)  ) 
    }
 
 	jsonBody, err := json.Marshal(content)
	if err != nil {
		http.Error(w, "Error converting content to json",
			http.StatusInternalServerError)
	}
	w.Write(jsonBody)
}
