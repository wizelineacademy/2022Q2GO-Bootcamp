package main

import (
	"fmt"
	"sync"
	"time"
	"github.com/waltergomezda/2022Q2GO-Bootcamp/internal/workerpool"
	"github.com/waltergomezda/2022Q2GO-Bootcamp/internal/service"
	"github.com/waltergomezda/2022Q2GO-Bootcamp/internal/repository"
)
type testTask struct {
	Name string
	Line int
	TaskProcessor func(int, string, int)
}
func (t testTask) Run(workerID int){
	t.TaskProcessor(workerID, t.Name, t.Line)
}

func main() {
	itemType := "even"
	totalItems := 10
	totalItemsPerWorker :=2

	//Reading CSV
	repo:= repository.NewCsvRepository("data.csv")
	svc := service.NewCsvService(repo)
	results, err := svc.ExtractCsvRecords()
	if err != nil {
		// http.Error(w, "Error retrieving CSV data",
		// 	http.StatusInternalServerError)
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

	sampleStringTaskFn := func(workerID int, name string, line int){
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
		resultsPerWorker[workerID] = fmt.Sprint(resultsPerWorker[workerID], "\n Line: ", line ," key:", tempResult.Key , " value: ", tempResult.Value)
	
		if name != ""{
			fmt.Println( fmt.Sprint("worker", workerID ," finished ", name))
		}
		wg.Done()
    }
    var tasks []testTask
	validTaskSize :=0
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
            TaskProcessor: sampleStringTaskFn,
        })
		validTaskSize++
    }
	wg.Add(validTaskSize)
    for _, task := range tasks{
        pool.ScheduleWork(task)
    }
    pool.Close()
    wg.Wait()

	for workerId, value := range resultsPerWorker {
        fmt.Println("Worker",workerId, ":", value)
    }
}