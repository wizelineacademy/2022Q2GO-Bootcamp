package main

import (
	"fmt"
	"github.com/waltergomezda/2022Q2GO-Bootcamp/internal/service"
	"github.com/waltergomezda/2022Q2GO-Bootcamp/internal/repository"
)
func main() {
	itemType := "even"
	maxItems := 5
	maxItemsPerWorker :=1
	numberOfWorkers := 7
	processedItems :=0

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
	
	for a := 1; a <= processedItems; a++ {
        fmt.Println(<-results)
    } 
}
