package pokemon

import (
	"context"
	"fmt"
	"sync"
)

type workFunc func(ctx context.Context, pokemon Pokemon) (Pokemon, error)

type workerPool struct {
	queue     chan Work
	result    chan Result
	done      chan any
	condition int
}

type Work struct {
	fn      workFunc
	pokemon Pokemon
}

type Result struct {
	Value Pokemon
	err   error
}

type workerInfo struct {
	id, limit, condition int
}

func NewWorkerPool(tpe string) workerPool {
	var condition int
	if tpe == "odd" {
		condition = 1
	}
	return workerPool{
		queue:     make(chan Work),
		result:    make(chan Result),
		done:      make(chan any),
		condition: condition,
	}
}

func worker(ctx context.Context, wInf workerInfo, wg *sync.WaitGroup, works <-chan Work, result chan<- Result) {
	defer wg.Done()
	fmt.Println("worker", wInf.id, "started")
	count := 0
	for {
		select {
		case work, ok := <-works:
			if !ok {
				return
			}
			if work.pokemon.ID%2 == wInf.condition {
				count++
				result <- Result{Value: work.pokemon, err: nil}
				fmt.Println("worker", wInf.id, "processed successfully", count)
			}
			if count == wInf.limit {
				return
			}
		case <-ctx.Done():
			fmt.Println("workers done")
			result <- Result{err: ctx.Err()}
			return
		}
	}
}

func (w Work) execute(ctx context.Context) Result {
	val, err := w.fn(ctx, w.pokemon)
	if err != nil {
		return Result{err: err}
	}
	return Result{Value: val}
}

func (w workerPool) Run(ctx context.Context, limit int) {
	var wg sync.WaitGroup

	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go worker(ctx, workerInfo{
			id:        i,
			limit:     limit,
			condition: w.condition,
		}, &wg, w.queue, w.result)
	}
	wg.Wait()
	w.Close()
}

func (w workerPool) Close() {
	close(w.done)
	close(w.result)
}

func (w workerPool) AddWork(works []Work) {
	for i := range works {
		w.queue <- works[i]
	}
	close(w.queue)
}
