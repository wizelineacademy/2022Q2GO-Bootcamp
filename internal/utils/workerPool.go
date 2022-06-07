package utils

// The major part of the code here was obtained from https://medium.com/code-chasm/go-concurrency-pattern-worker-pool-a437117025b1

// T is a type alias to accept any type.
type T = interface{}

// WorkerPool is a contract for Worker Pool implementation
type WorkerPool interface {
	Run()
	AddTask(task func())
}

type workerPool struct {
	maxWorker   int
	queuedTaskC chan func()
}

func (wp *workerPool) Run() {
	for i := 0; i < wp.maxWorker; i++ {
		go func(workerID int) {
			for task := range wp.queuedTaskC {
				task()
			}
		}(i + 1)
	}
}

func (wp *workerPool) AddTask(task func()) {
	wp.queuedTaskC <- task
}

func NewWorkerPool(totalWorkers int) WorkerPool {
	return &workerPool{
		maxWorker:   totalWorkers,
		queuedTaskC: make(chan func()),
	}
}
