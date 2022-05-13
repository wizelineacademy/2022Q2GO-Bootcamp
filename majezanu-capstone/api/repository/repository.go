package repository

import (
	"encoding/csv"
	"errors"
	"io"
	"majezanu/capstone/api/models"
	"os"
	"strconv"
)

const dataFilePath = "./data/pokemon.csv"

func FetchPokemonData() (data []models.Pokemon, err error) {
	csvFile, err := os.Open(dataFilePath)
	if err != nil {
		return
	}
	defer csvFile.Close()
	csvLines, err := csv.NewReader(csvFile).ReadAll()
	if err != nil {
		return
	}
	for _, line := range csvLines {
		id, err := strconv.Atoi(line[0])
		if err != nil {
			break
		}
		item := models.Pokemon{
			Id:   id,
			Name: line[1],
		}
		data = append(data, item)
	}
	return
}

func FetchPokemonDataById(idToSearch int) (data models.Pokemon, err error) {
	csvFile, err := os.Open(dataFilePath)
	if err != nil {
		return
	}
	defer csvFile.Close()

	csvReader := csv.NewReader(csvFile)

	for {
		line, err := csvReader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			break
		}
		id, err := strconv.Atoi(line[0])
		if err != nil {
			break
		}
		if id == idToSearch {
			data = models.Pokemon{
				Id:   id,
				Name: line[1],
			}
			break
		}
	}

	if data.Id != idToSearch {
		err = errors.New("There is no pokemon with id")
	}

	return
}

func FetchPokemonDataByName(nameToSearch string) (data models.Pokemon, err error) {
	csvFile, err := os.Open(dataFilePath)
	if err != nil {
		return
	}
	defer csvFile.Close()

	csvReader := csv.NewReader(csvFile)

	for {
		line, err := csvReader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			break
		}
		id, err := strconv.Atoi(line[0])
		if err != nil {
			break
		}

		if nameToSearch == line[1] {
			data = models.Pokemon{
				Id:   id,
				Name: line[1],
			}
			break
		}
	}

	if data.Name != nameToSearch {
		err = errors.New("There is no pokemon with name")
	}

	return
}
