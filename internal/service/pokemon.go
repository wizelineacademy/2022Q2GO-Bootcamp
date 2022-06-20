package service

import (
	"context"
	"fmt"
	"sync"

	"github.com/krmirandas/2022Q2GO-Bootcamp/internal/entity"
)

// PokemonService the contract of the pokemon service
type PokemonService interface {
	// CreatePokemon create new record Pokemon
	CreatePokemon(ctx context.Context, Pokemon *entity.Pokemon) error

	// FindPokemon gets filtered specific Pokemon
	FindPokemon(ctx context.Context, filter *entity.Pokemon) ([]entity.Pokemon, error)

	Count(ctx context.Context) (int, error)
}

type PokemonRepo interface {
	ReadPokemon() ([]entity.Pokemon, error)
	WritePokemon(pokemon *entity.Pokemon) error
	Count(ctx context.Context) (int, error)
}

type pokemonService struct {
	repo PokemonRepo
}

func NewPokemonService(repo PokemonRepo) PokemonService {
	return &pokemonService{repo: repo}
}

func (ps *pokemonService) CreatePokemon(ctx context.Context, Pokemon *entity.Pokemon) error {
	return nil
}

func (ps *pokemonService) FindPokemon(ctx context.Context, filter *entity.Pokemon) ([]entity.Pokemon, error) {
	pokemons, err := ps.repo.ReadPokemon(ctx)
	if err != nil {
		return nil, err
	}

	return pokemons, nil
}

func (ps *pokemonService) Count(ctx context.Context) (int, error) {
	return ps.repo.Count(ctx)
}

var mu sync.Mutex

// with Worker pools
// func ConcuRSwWP(f *os.File, typeP string, items int, itemsPerWorker int) []model.Pokemon {
func ConcuRSwWP(f *os.File, itemsPerWorker int) []entity.Pokemon {
	fcsv := csv.NewReader(f)
	rs := make([]model.Pokemon, 0)
	numWps := itemsPerWorker
	jobs := make(chan []string, numWps)
	res := make(chan *model.Pokemon)

	var wg sync.WaitGroup
	worker := func(jobs <-chan []string, results chan<- *entity.Pokemon) {
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

func ParseStruct(data []string) *entity.Pokemon {

	return &entity.Pokemon{
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
