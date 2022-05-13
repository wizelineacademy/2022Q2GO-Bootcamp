package database

import (
	"context"
	"encoding/csv"
	"errors"
	"log"
	"os"
	"strconv"

	"github.com/TanZng/toh-api/models"
)

type CVSRepository struct {
	file string
}

func NewCSVRepository(file string) (*CVSRepository, error) {

	return &CVSRepository{file}, nil
}

func (repo *CVSRepository) InsertCharacter(ctx context.Context, user *models.Character) error {

	return nil
}

func (repo *CVSRepository) GetCharacterById(ctx context.Context, id int64) (*models.Character, error) {
	f, err := os.Open(repo.file)
	if err != nil {
		log.Fatal("Unable to read input file "+repo.file, err)
	}
	defer f.Close()

	records, err := csv.NewReader(f).ReadAll()
	if err != nil {
		log.Println("Error READ CSV:", err)
		return nil, err
	}

	// log.Println("Records:", records)

	var find bool
	var character models.Character

	for _, record := range records {
		readId, _ := strconv.ParseInt(record[0], 10, 64)
		if readId == id {
			character.ID = readId
			character.Age, _ = strconv.ParseInt(record[2], 10, 64)
			character.Name = record[1]
			find = true
			break
		}
	}
	if !find {
		err := errors.New("CSV: Character Not found")
		return nil, err
	}
	return &character, nil
}

// func (repo *CVSRepository) GetAllCharacter(ctx context.Context, id int64) ([]models.Character, error) {
// 	var employeeRecords []models.Character
// 	records, err := csv.NewReader(repo.file).ReadAll()
// 	if err != nil {
// 		return employeeRecords, err
// 	}
// 	for _, record := range records {
// 		idInt, _ := strconv.ParseInt(record[0], 10, 64)
// 		ageInt, _ := strconv.ParseUint(record[2], 10, 32)
// 		data := models.Character{
// 			ID:   idInt,
// 			Name: record[1],
// 			Age:  ageInt,
// 		}
// 		employeeRecords = append(employeeRecords, data)
// 	}

// 	return employeeRecords, nil
// }

func (repo *CVSRepository) Close() error {
	log.Println("Close CSV")
	return nil
}
