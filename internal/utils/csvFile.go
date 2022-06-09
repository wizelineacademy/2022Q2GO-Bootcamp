package utils

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/McAdam17/2022Q2GO-Bootcamp/internal/entity"
)

func ReadPokemonDataFromCSVFile(fileName string) ([]entity.Pokemon, map[int]string, map[string]int, error) {
	if !strings.HasSuffix(fileName, ".csv") {
		return nil, nil, nil, fmt.Errorf("invalid file type")
	}

	csvFile, err := os.Open(fileName)
	if err != nil {
		return nil, nil, nil, err
	}

	defer csvFile.Close()

	csvRows, err := csv.NewReader(csvFile).ReadAll()
	if err != nil {
		return nil, nil, nil, fmt.Errorf("invalid file type")
	}

	pokemons := []entity.Pokemon{}
	pokemonsByID := make(map[int]string)
	pokemonsByName := make(map[string]int)
	for _, row := range csvRows {
		id, err := strconv.Atoi(row[0])
		if err != nil {
			return nil, nil, nil, fmt.Errorf("invalid file type")
		}
		name := row[1]
		name = strings.ToLower(name)

		pokemon := entity.Pokemon{
			ID:   id,
			Name: name,
		}

		pokemons = append(pokemons, pokemon)
		pokemonsByID[id] = name
		pokemonsByName[name] = id
	}

	return pokemons, pokemonsByID, pokemonsByName, nil
}

func FindPokemonDataFromCSVFile(fileName string, pokemonID int) (*entity.Pokemon, error) {
	if !strings.HasSuffix(fileName, ".csv") {
		return nil, fmt.Errorf("invalid file type")
	}

	csvFile, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}

	defer csvFile.Close()

	csvRows, err := csv.NewReader(csvFile).ReadAll()
	if err != nil {
		return nil, fmt.Errorf("invalid file type")
	}

	pokemon := &entity.Pokemon{}
	for _, row := range csvRows {
		id, err := strconv.Atoi(row[0])
		if err != nil {
			return nil, fmt.Errorf("invalid file type")
		}
		name := row[1]
		name = strings.ToLower(name)

		if id == pokemonID {
			pokemon.ID = id
			pokemon.Name = name
		}
	}

	if len(pokemon.Name) == 0 {
		return nil, fmt.Errorf("pokemon not found")
	}

	return pokemon, nil
}

func WritePokemonsOnCSV(pokemons []entity.Pokemon, fileName string) error {
	csvFile, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		return err
	}

	defer csvFile.Close()

	for _, pokemon := range pokemons {
		pokemonPlainText := fmt.Sprintf("\n%d,%s", pokemon.ID, pokemon.Name)

		if _, err = csvFile.WriteString(pokemonPlainText); err != nil {
			return err
		}
	}

	return nil
}
