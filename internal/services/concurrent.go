package services

import (
	"fmt"
	"sync"
)

var (
	// list of channels to communicate with workers
	// workers accessed synchronousely no mutex required
	workers = make(map[string]chan []string)

	// wg is to make sure all workers done before exiting main
	wg = sync.WaitGroup{}

	// mu used only for sequential printing, not relevant for program logic
	mu = sync.Mutex{}
)

func process(rec []string) {
	l := len(rec)
	part := rec[l-1]

	if c, ok := workers[part]; ok {
		// send rec to worker
		c <- rec
	} else {
		// if no worker for the partition

		// make a chan
		nc := make(chan []string)
		workers[part] = nc

		// start worker with this chan
		go worker(nc)

		// send rec to worker via chan
		nc <- rec
	}
}

func worker(c chan []string) {

	// wg.Done signals to main worker completion
	wg.Add(1)
	defer wg.Done()

	part := [][]string{}
	for {
		// wait for a rec or close(chan)
		rec, ok := <-c
		if ok {
			// save the rec
			// instead of accumulation in memory
			// this can be saved to file directly
			part = append(part, rec)
		} else {
			// channel closed on EOF

			// dump partition
			// locks ensures sequential printing
			// not a required for independent files
			mu.Lock()
			for _, p := range part {
				fmt.Printf("%+v\n", p)
			}
			mu.Unlock()

			return
		}
	}

}

// simply signals to workers to stop
func savePartitions() {
	for _, c := range workers {
		// signal to all workers to exit
		close(c)
	}
}
