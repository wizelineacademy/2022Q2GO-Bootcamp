package repository

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"sync"

	"github.com/krmirandas/2022Q2GO-Bootcamp/internal/model"
)

var mu sync.Mutex

// with Worker pools
// func ConcuRSwWP(f *os.File, typeP string, items int, itemsPerWorker int) []model.Pokemon {
func ConcuRSwWP(f *os.File, itemsPerWorker int) []model.Pokemon {
	fcsv := csv.NewReader(f)
	rs := make([]model.Pokemon, 0)
	numWps := itemsPerWorker
	jobs := make(chan []string, numWps)
	res := make(chan *model.Pokemon)

	var wg sync.WaitGroup
	worker := func(jobs <-chan []string, results chan<- *model.Pokemon) {
		for {
			select {
			case job, ok := <-jobs:
				if !ok {
					return
				}
				results <- ParseStruct(job)
			}
		}
	}

	// init workers
	for w := 0; w < numWps; w++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			worker(jobs, res)
		}()
	}

	go func() {
		for w := 0; w < numWps; w++ {
			rStr, err := fcsv.Read()
			if err == io.EOF {
				break
			}
			if err != nil {
				fmt.Println("ERROR: ", err.Error())
				break
			}
			jobs <- rStr
		}
		close(jobs) // close jobs to signal workers that no more job are incoming.
	}()

	go func() {
		wg.Wait()
		close(res) // when you close(res) it breaks the below loop.
	}()

	for r := range res {
		// fmt.Println(*r)
		rs = append(rs, *r)
	}

	return rs

}

func ParseStruct(data []string) *model.Pokemon {

	return &model.Pokemon{
		ID:         data[0],
		Name:       data[1],
		Type1:      data[2],
		Type2:      data[3],
		Total:      data[4],
		HP:         data[5],
		Attack:     data[6],
		Defense:    data[7],
		SpAtk:      data[8],
		SpDef:      data[9],
		Speed:      data[10],
		Generation: data[11],
		Legendary:  data[12],
	}
}
