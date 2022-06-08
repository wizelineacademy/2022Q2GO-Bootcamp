package repository

import (
	"fmt"

	"github.com/McAdam17/2022Q2GO-Bootcamp/internal/entity"
	"github.com/McAdam17/2022Q2GO-Bootcamp/internal/utils"
)

type pokemonRepository struct {
	fileNameStore   string
	readData        []entity.Pokemon
	queryDataByID   map[int]string
	queryDataByName map[string]int
}

func NewPokemonRepository(fileName string) *pokemonRepository {
	arrayData, mapDataByID, mapDataByName, err := utils.ReadPokemonDataFromCSVFile(fileName)
	if err != nil {
		panic("Error reading file " + err.Error())
	}

	return &pokemonRepository{
		fileNameStore:   fileName,
		readData:        arrayData,
		queryDataByID:   mapDataByID,
		queryDataByName: mapDataByName,
	}
}

func (pR *pokemonRepository) GetAllPokemons() ([]entity.Pokemon, error) {
	return pR.readData, nil
}

func (pR *pokemonRepository) GetPokemonById(id int) (*entity.Pokemon, error) {
	pokemonName, ok := pR.queryDataByID[id]
	if !ok {
		return nil, fmt.Errorf("pokemon not found")
	}

	return &entity.Pokemon{
		ID:   id,
		Name: pokemonName,
	}, nil
}

func (pR *pokemonRepository) GetPokemonItemsFromCSV(typeReading string, items, itemsPerWorkers int) ([]entity.Pokemon, error) {
	pokemons := make([]entity.Pokemon, items)

	currentDataLen := len(pR.readData)

	id := 1
	if typeReading == "even" {
		id = 2
	}

	index := 0
	for (index < items) && (id < currentDataLen) {
		err := addPokemonWithIDFormCSVFileToArrayAtIndex(id, pR.fileNameStore, pokemons, index)
		if err != nil {

		}

		index++
		id += 2
	}

	pokemons = utils.CleanPokemonsResponse(pokemons)

	return pokemons, nil
}

func addPokemonWithIDFormCSVFileToArrayAtIndex(id int, fileName string, pokemons []entity.Pokemon, index int) error {
	pokemon, err := utils.FindPokemonDataFromCSVFile(fileName, id)
	if err != nil {
		return err
	}

	pokemons[index] = *pokemon
	return nil
}

func (pR *pokemonRepository) AddNewPokemons(newPokemons []string) ([]entity.Pokemon, error) {
	startIndex := -1

	for _, pokemonName := range newPokemons {
		_, pokemonAlreadyExists := pR.queryDataByName[pokemonName]
		if pokemonAlreadyExists {
			continue
		}

		if startIndex == -1 {
			startIndex = len(pR.readData)
		}

		newPokemonID := len(pR.readData) + 1
		newPokemon := &entity.Pokemon{
			ID:   newPokemonID,
			Name: pokemonName,
		}

		pR.addNewPokemonToMemoryData(newPokemon)
	}

	if startIndex == -1 {
		return pR.readData, nil
	}

	err := pR.addNewPokemonsToCSV(startIndex)
	if err != nil {
		return nil, err
	}

	return pR.readData, nil
}

func (pR *pokemonRepository) addNewPokemonToMemoryData(newPokemon *entity.Pokemon) {
	pR.queryDataByID[newPokemon.ID] = newPokemon.Name
	pR.queryDataByName[newPokemon.Name] = newPokemon.ID

	pR.readData = append(pR.readData, *newPokemon)
}

func (pR *pokemonRepository) addNewPokemonsToCSV(index int) error {
	currentLen := len(pR.readData)
	partialLen := currentLen - index
	partialPokemonData := make([]entity.Pokemon, partialLen)
	copy(partialPokemonData[:], pR.readData[index:])

	return utils.WritePokemonsOnCSV(partialPokemonData, pR.fileNameStore)
}
