package workerpool

import (
	// "fmt"
	"sync"
)

type WorkFunc interface {
	Run(int)
}

type GoroutinePool struct {
	queue chan work
	wg sync.WaitGroup
}

type work struct {
	fn WorkFunc
}
// NewGorountinePool creates a new pool of gorutines to schedule async work.
func NewGoroutinePool(workerSize int) *GoroutinePool {
	gp := &GoroutinePool{
		queue: make(chan work),
	}
	gp.AddWorkers(workerSize)
	return gp
}
//Close waits for all goroutines to terminate
func (gp *GoroutinePool) Close(){
	close(gp.queue)
	gp.wg.Wait()
}
//ScheduleWork registers the given functino to be executed at some point. The given param will
// be supplied to the function during execution.
func (gp *GoroutinePool) ScheduleWork(fn WorkFunc){
	gp.queue <- work{fn}
}
// AddWorkers introduces more goroutines in the worker pool, increasing potential parallelism.
func (gp *GoroutinePool) AddWorkers(numWorkers int){
	gp.wg.Add(numWorkers)
	for i:=0; i< numWorkers; i++ {
		go func(workerID int){
			count := 0
			for job := range gp.queue {
				job.fn.Run(workerID)
				count++
			}
			// fmt.Println("worker ",workerID," executed ", count, " tasks")
			gp.wg.Done()
		}(i)
	}
}