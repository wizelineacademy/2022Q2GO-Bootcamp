package repository

import (
	"fmt"
	"sync"
	"toh-api/internal/entity"
)

//type WorkFunc interface {
//	Run() interface{}
//}

type Job interface {
	getId() int
	getTask() interface{}
	RunTask() interface{}
}

type Result struct {
	job    Job
	result interface{}
}

func (gp *GoroutinePool) Worker(wg *sync.WaitGroup) {
	for job := range gp.jobs {
		output := Result{job, job.RunTask()}
		gp.results <- output
	}
	wg.Done()
}

type GoroutinePool struct {
	function    func()
	jobs        chan Job
	results     chan Result
	ResultsPool []interface{}
	workerSize  int
}

// NewGoroutinePool creates a new pool of goroutines to schedule async work
func NewGoroutinePool(workerSize int) *GoroutinePool {
	gp := &GoroutinePool{
		workerSize: workerSize,
		jobs:       make(chan Job),
		results:    make(chan Result),
	}
	return gp
}

func (gp *GoroutinePool) AddWorkers() {
	var wg sync.WaitGroup
	for i := 0; i < gp.workerSize; i++ {
		wg.Add(1)
		go gp.Worker(&wg)
	}
	wg.Wait()
	close(gp.results)
}

func (gp *GoroutinePool) AllocateJobs(jobs []Job) {
	for _, job := range jobs {
		gp.jobs <- job
	}

	close(gp.jobs)
}

func (gp *GoroutinePool) GetResult(done chan bool) {
	for result := range gp.results {
		fmt.Printf("Job id %d, input task %+v , result %+v\n", result.job.getId(), result.job.getTask(), result.result)
		gp.ResultsPool = append(gp.ResultsPool, result.result)
	}
	done <- true
}

func (gp *GoroutinePool) ObtainCharacters() []entity.Character {
	var characters []entity.Character
	for _, r := range gp.ResultsPool {
		characters = append(characters, r.(entity.Character))
	}
	return characters
}
