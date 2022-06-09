package pokemon

import (
	"context"
	"encoding/json"
	"fmt"
	errs "github.com/esvarez/go-api/pkg/error"
	"log"
	"net/http"
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

func NewService(repo repository) *Service {
	return &Service{
		endpoint: endpoint,
		repo:     repo,
		client:   http.Client{},
	}
}

func (s Service) FindByID(id string) (*Pokemon, error) {
	resp, err := s.client.Get(s.endpoint + pokemonInfo + id)
	if err != nil {
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

	pokemons, err := s.repo.GetAllPokemon()
	if err != nil {
		return nil, err
	}

	for _, pokemon := range pokemons {
		tasks = append(tasks, Work{
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
				fmt.Println("Reached EOF")
				continue
			}

			pkmns = append(pkmns, r.Value)
			if len(pkmns) == items {
				fmt.Println("Get pokemon requested")
				return pkmns, nil
			}
		case <-pool.done:
			return pkmns, nil
		}
	}
}
