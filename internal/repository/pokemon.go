package repository

import (
	"bufio"
	"encoding/csv"
	"errors"
	"io"
	"os"

	"github.com/krmirandas/2022Q2GO-Bootcamp/internal/entity"
)

type PokemonRepo interface {

	// ReadPokemon reads and parse all pokemon records from
	// the data csv file
	ReadPokemon(offset, limit int) ([]entity.Pokemon, error)

	ReadOnePokemon(id string) (entity.Pokemon, error)

	// Count returns the number of albums.
	Count() (int, error)
}

// pokemonRepo the pokemon repository implementation struct
type pokemonRepo struct {
	filePath string
}

func NewPokemonRepo(file string) PokemonRepo {
	return &pokemonRepo{filePath: file}
}

// IMPLEMENTATION -------------------

/************************** FIRST DELIVERY ************************/
func (pr *pokemonRepo) ReadOnePokemon(id string) (entity.Pokemon, error) {
	var pokemon entity.Pokemon

	//ID,Name,Type1,Type2,Total,HP,Attack,Defense,SpAtk,SpDef,Speed,Generation,Legendary
	csvFile, _ := os.Open(pr.filePath)
	reader := csv.NewReader(bufio.NewReader(csvFile))
	for {
		line, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			return pokemon, err
		}

		if line[0] == id {
			pokemon = entity.Pokemon{
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
			}
		}
	}

	if pokemon.ID == "" {
		return pokemon, errors.New("empty struct")
	}

	return pokemon, nil
}

func (pr *pokemonRepo) ReadPokemon(offset, limit int) ([]entity.Pokemon, error) {
	var pokemons []entity.Pokemon

	//ID,Name,Type1,Type2,Total,HP,Attack,Defense,SpAtk,SpDef,Speed,Generation,Legendary
	csvFile, _ := os.Open(pr.filePath)
	reader := csv.NewReader(bufio.NewReader(csvFile))
	for {
		line, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			return nil, err
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

	return paginate(pokemons, offset, limit), nil
}

// Count returns the number of the album records in the database
func (pr *pokemonRepo) Count() (int, error) {
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

/************************** FIRST DELIVERY ************************/

func paginate(x []entity.Pokemon, skip int, size int) []entity.Pokemon {
	if skip > len(x) {
		skip = len(x)
	}

	end := skip + size
	if end > len(x) {
		end = len(x)
	}

	return x[skip:end]
}
