package repository

import (
	"bufio"
	"context"
	"encoding/csv"
	"io"
	"log"
	"os"

	"github.com/krmirandas/2022Q2GO-Bootcamp/internal/entity"
)

type PokemonRepo interface {

	// ReadPokemon reads and parse all pokemon records from
	// the data csv file
	ReadPokemon(ctx context.Context) ([]entity.Pokemon, error)

	// WritePokemon writes new record to the data csv file
	WritePokemon(ctx context.Context, pokemon *entity.Pokemon) error

	// Count returns the number of albums.
	Count(ctx context.Context) (int, error)
}

// pokemonRepo the pokemon repository implementation struct
type pokemonRepo struct {
	filePath string
}

func NewPokemonRepo(file string) PokemonRepo {
	return &pokemonRepo{filePath: file}
}

// IMPLEMENTATION -------------------

func (pr *pokemonRepo) ReadPokemon(ctx context.Context) ([]entity.Pokemon, error) {
	log.Println()
	var pokemons []entity.Pokemon

	//ID,Name,Type1,Type2,Total,HP,Attack,Defense,SpAtk,SpDef,Speed,Generation,Legendary
	csvFile, _ := os.Open(pr.filePath)
	reader := csv.NewReader(bufio.NewReader(csvFile))
	for {
		line, error := reader.Read()
		if error == io.EOF {
			break
		} else if error != nil {
			log.Fatal(error)
		}

		pokemons = append(pokemons, entity.Pokemon{
			ID:         line[0],
			Name:       line[1],
			Type1:      line[2],
			Type2:      line[3],
			Total:      line[4],
			HP:         line[5],
			Attack:     line[6],
			Defense:    line[7],
			SpAtk:      line[8],
			SpDef:      line[9],
			Speed:      line[10],
			Generation: line[11],
			Legendary:  line[12],
		})
	}

	return pokemons, nil
}

func (*pokemonRepo) WritePokemon(ctx context.Context, pokemon *entity.Pokemon) error {
	log.Println()
	return nil
}

// Count returns the number of the album records in the database.
func (pr *pokemonRepo) Count(ctx context.Context) (int, error) {
	openfile, err := os.Open(pr.filePath)
	if err != nil {
		return 0, err
	}
	filedata, err := csv.NewReader(openfile).ReadAll()
	if err != nil {
		return 0, err
	}
	totalQuestions := len(filedata)
	return totalQuestions, nil
}
