package utils

import "github.com/McAdam17/2022Q2GO-Bootcamp/internal/entity"

// The major part of the code here was obtained from https://medium.com/code-chasm/go-concurrency-pattern-worker-pool-a437117025b1

// T is a type alias to accept any type.
type T = interface{}

type Result struct {
	Err error
}

type Task struct {
	Index       int
	ID          int
	Pokemons    []entity.Pokemon
	FileName    string
	ResultsChan chan Result
}

// WorkerPool is a contract for Worker Pool implementation
type WorkerPool interface {
	Run()
	AddTask(task Task)
}

type workerPool struct {
	maxWorker   int
	queuedTaskC chan Task
}

func (wp *workerPool) Run() {
	for i := 0; i < wp.maxWorker; i++ {
		go func(workerID int) {
			for task := range wp.queuedTaskC {
				pokemon, err := FindPokemonDataFromCSVFile(task.FileName, task.ID)
				task.Pokemons[task.Index] = *pokemon

				task.ResultsChan <- Result{
					Err: err,
				}
			}
		}(i + 1)
	}
}

func (wp *workerPool) AddTask(task Task) {
	wp.queuedTaskC <- task
}

func NewWorkerPool(totalWorkers int) WorkerPool {
	return &workerPool{
		maxWorker:   totalWorkers,
		queuedTaskC: make(chan Task),
	}
}
