package usecase

import (
	"encoding/csv"
	"fmt"
	"strconv"

	"github.com/luischitala/2022Q2GO-Bootcamp/models"
	"github.com/luischitala/2022Q2GO-Bootcamp/repository"
)

//Interface that allows to execute all the entity operations
type Character interface {
	Reader
}
type Reader interface {
	ReadCsv() ([]models.Character, error)
}

//Struct that separate the logic between the repository and the superior layers
type rs struct {
	Csv repository.Csv
}

func ReadCsv(r *rs) ([]models.Character, error) {
	characters := make([]models.Character, 0)
	csvFile, err := r.Csv.ReadCsvFile()
	if err != nil {
		fmt.Println(err)
	}
	csvLines, err := csv.NewReader(csvFile).ReadAll()

	if err != nil {
		fmt.Println(err)
	}
	for _, line := range csvLines {
		id, _ := strconv.Atoi(line[0])

		character := models.Character{
			Id:   id,
			Name: line[1],
		}
		characters = append(characters, character)

	}
	return characters, nil
}
