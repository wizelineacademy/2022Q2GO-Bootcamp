package pokemon

import (
	"context"
	"encoding/json"
	"fmt"
	errs "github.com/esvarez/go-api/pkg/error"
	"log"
	"net/http"
	"sync"
)

const (
	endpoint    = "https://pokeapi.co/api/v2/"
	pokemonInfo = "pokemon/"

	numWorkers = 8
)

type writer interface {
	AddPokemon(pokemon *Pokemon) error
}

type reader interface {
	GetAllPokemon() ([]Pokemon, error)
}

type repository interface {
	writer
	reader
}

type Service struct {
	endpoint string
	workerPool
	client http.Client
	repo   repository
}

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

func (w Work) execute(ctx context.Context) Result {
	val, err := w.fn(ctx, w.pokemon)
	if err != nil {
		return Result{err: err}
	}
	return Result{Value: val}
}

func worker(ctx context.Context, id, limit, condition int, wg *sync.WaitGroup, works <-chan Work, result chan<- Result) {
	defer wg.Done()
	fmt.Println("worker", id, "started")
	count := 0
	for {
		select {
		case work, ok := <-works:
			if !ok {
				return
			}

			// TODO handle error
			if work.pokemon.ID%2 == condition {
				count++
				result <- Result{Value: work.pokemon, err: nil}
			}
			fmt.Println("worker", id, "started", "processed", count)
			if count == limit {
				return
			}
		case <-ctx.Done():
			fmt.Println("workers done")
			result <- Result{err: ctx.Err()}
			return
		}
	}
}

func NewWorkerPool(tpe string) workerPool {
	// TODO move worker pool to other package
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

func (w workerPool) Run(ctx context.Context, limit int) {
	var wg sync.WaitGroup

	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go worker(ctx, i, limit, w.condition, &wg, w.queue, w.result)
	}
	wg.Wait()
	w.Close()
}

func NewService(repo repository) *Service {
	return &Service{
		endpoint: endpoint,
		repo:     repo,
		client:   http.Client{},
		//workerPool: NewWorkerPool(),
	}
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

func (s Service) FindByID(id string) (*Pokemon, error) {
	resp, err := s.client.Get(s.endpoint + pokemonInfo + id)
	if err != nil {
		log.Println("error getting pokemon: ", err)
		return nil, err
	}
	if resp.StatusCode == http.StatusNotFound {
		return nil, errs.ErrNotFound
	}

	pokemon := &Pokemon{}
	json.NewDecoder(resp.Body).Decode(pokemon)

	if err := s.repo.AddPokemon(pokemon); err != nil {
		log.Println("error adding pokemon: ", err)
		return nil, err
	}

	return pokemon, nil
}

func (s Service) GetPokemon(tpe string, items, itemsWorker int) ([]Pokemon, error) {
	pool := NewWorkerPool(tpe)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var tasks []Work

	/*
		simpleTask := func(ctx context.Context, pokemon Pokemon) (Pokemon, error) {
			fmt.Println("Validating ", pokemon)
			return pokemon, nil
		}
	*/
	pokemons, err := s.repo.GetAllPokemon()
	if err != nil {
		return nil, err
	}

	for _, pokemon := range pokemons {
		tasks = append(tasks, Work{
			//fn:      simpleTask,
			pokemon: pokemon,
		})
	}
	go pool.AddWork(tasks)

	go pool.Run(ctx, itemsWorker)

	pkmns := []Pokemon{}
	for {
		select {
		case r, ok := <-pool.result:
			if !ok {
				continue
			}

			pkmns = append(pkmns, r.Value)
			if len(pkmns) == items {
				return pkmns, nil
			}
		case <-pool.done:
			return pkmns, nil
		}
	}
}
