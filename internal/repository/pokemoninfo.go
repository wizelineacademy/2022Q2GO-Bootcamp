package repository

import (
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"sync"

	"github.com/krmirandas/2022Q2GO-Bootcamp/internal/entity"
)

type PokemonInfoRepo interface {
	// WritePokemon writes new record to the data csv file
	WritePokemon(pokemon entity.PokemonInfo) error
	ConcuRSwWP(itemsPerWorker int) []entity.PokemonInfo
}

// pokemonRepo the pokemon repository implementation struct
type pokemonInfoRepo struct {
	filePath string
}

func NewPokemonInfoRepo(file string) PokemonInfoRepo {
	return &pokemonInfoRepo{filePath: file}
}

/************************** SECOND DELIVERY ************************/

func (pir *pokemonInfoRepo) WritePokemon(pokemon entity.PokemonInfo) error {
	var csvFile *os.File
	var err error

	fileName := pir.filePath
	if _, err := os.Stat(fileName); err == nil {
		// path/to/whatever exists
		csvFile, err = os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY, 0644)
	} else if errors.Is(err, os.ErrNotExist) {
		// path/to/whatever does *not* exist
		csvFile, err = os.Create(fmt.Sprintf(fileName))
	}

	if err != nil {
		log.Fatalf("failed creating file: %s", err)
		return err
	}

	w := csv.NewWriter(csvFile)
	defer w.Flush()

	// Using Write
	fmt.Println(pokemon)
	row := []string{strconv.Itoa(pokemon.ID), pokemon.Name, strconv.Itoa(pokemon.BaseExperience),
		strconv.Itoa(pokemon.Height), strconv.FormatBool(pokemon.IsDefault), strconv.Itoa(pokemon.Order),
		strconv.Itoa(pokemon.Weight), pokemon.LocationAreaEncounters}
	if err := w.Write(row); err != nil {
		// log.Fatalln("error writing record to file", err)
		return err
	}

	return nil
}

// with Worker pools
// func ConcuRSwWP(f *os.File, typeP string, items int, itemsPerWorker int) []entity.Pokemon {
func (pir *pokemonInfoRepo) ConcuRSwWP(itemsPerWorker int) []entity.PokemonInfo {
	f, _ := os.Open(pir.filePath)
	defer f.Close()

	fcsv := csv.NewReader(f)
	rs := make([]entity.PokemonInfo, 0)
	numWps := itemsPerWorker
	jobs := make(chan []string, numWps)
	res := make(chan *entity.PokemonInfo)

	var wg sync.WaitGroup
	worker := func(jobs <-chan []string, results chan<- *entity.PokemonInfo) {
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

func ParseStruct(data []string) *entity.PokemonInfo {

	id, _ := strconv.Atoi(data[0])
	baseExperience, _ := strconv.Atoi(data[2])
	height, _ := strconv.Atoi(data[3])
	isDefault, _ := strconv.ParseBool(data[4])
	order, _ := strconv.Atoi(data[5])
	weight, _ := strconv.Atoi(data[6])

	return &entity.PokemonInfo{
		ID:                     id,
		Name:                   data[1],
		BaseExperience:         baseExperience,
		Height:                 height,
		IsDefault:              isDefault,
		Order:                  order,
		Weight:                 weight,
		LocationAreaEncounters: data[7],
	}
}
