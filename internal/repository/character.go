package repository

import (
	"encoding/csv"
	"errors"
	"log"
	"os"
	"strconv"
	"toh-api/internal/entity"
)

// CharacterRepository the contract of movie repository
type CharacterRepository interface {
	InsertCharacter(character *entity.Character) error
	GetCharacterById(id int64) (*entity.Character, error)
}

// characterRepository the character repository implementation struct
type characterRepository struct {
	file string
}

func NewCharacterRepository(file string) CharacterRepository {
	return &characterRepository{file}
}

func (repo *characterRepository) InsertCharacter(character *entity.Character) error {

	return nil
}

func (repo *characterRepository) GetCharacterById(id int64) (*entity.Character, error) {
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
	var character entity.Character

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

// func (repo *characterRepository) GetAllCharacter(ctx context.Context, id int64) ([]entity.Character, error) {
// 	var employeeRecords []entity.Character
// 	records, err := csv.NewReader(repo.file).ReadAll()
// 	if err != nil {
// 		return employeeRecords, err
// 	}
// 	for _, record := range records {
// 		idInt, _ := strconv.ParseInt(record[0], 10, 64)
// 		ageInt, _ := strconv.ParseUint(record[2], 10, 32)
// 		data := entity.Character{
// 			ID:   idInt,
// 			Name: record[1],
// 			Age:  ageInt,
// 		}
// 		employeeRecords = append(employeeRecords, data)
// 	}

// 	return employeeRecords, nil
// }
