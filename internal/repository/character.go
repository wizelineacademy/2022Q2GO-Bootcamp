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
	GetAllCharacters(givenType string, items, itemsPerWorkers int) ([]entity.Character, error)
}

// characterRepository the character repository implementation struct
type characterRepository struct {
	file string
}

func NewCharacterRepository(file string) CharacterRepository {
	return &characterRepository{file}
}

// Implementation

// InsertCharacter adds a character record
func (repo *characterRepository) InsertCharacter(character *entity.Character) error {
	// TODO implement me
	panic("implement me")
}

// GetCharacterById returns the character record with the given id
func (repo *characterRepository) GetCharacterById(id int64) (*entity.Character, error) {

	records, err := repo.getRecords()
	if err != nil {
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

// characterJob is the struct to handle the jobs related to get characters for the worker pool
type characterJob struct {
	id        int
	task      []string
	givenType string
}

func (tk characterJob) getId() int {
	return tk.id
}

func (tk characterJob) getTask() interface{} {
	return tk.task
}

func (tk characterJob) RunTask() interface{} {
	// specify task

	if tk.givenType == "odd" {
		return getOddRecords(tk.task)
	} else {
		return getEvenRecords(tk.task)
	}
}

// GetAllCharacters concurrently using a workerpool
// type: Only support "odd" or "even"
// items: Is an Int and is the number of valid items you need to display as a response
// items_per_workers: Is an Int and is the number of valid items the worker should append to the response
func (repo *characterRepository) GetAllCharacters(givenType string, items, itemsPerWorkers int) ([]entity.Character, error) {

	records, _ := repo.getRecords()

	pool := NewGoroutinePool(items / itemsPerWorkers)
	taskSize := items

	// wait for jobs to finish
	//wg := &sync.WaitGroup{}
	//wg.Add(taskSize)
	//
	var tasks []Job
	for v := 0; v < taskSize; v++ {
		tasks = append(tasks, characterJob{
			id:   v,
			task: records[v],
		})
	}

	go pool.AllocateJobs(tasks)

	done := make(chan bool)

	go pool.GetResult(done)

	pool.AddWorkers()
	<-done
	return pool.ObtainCharacters(), nil
}

func getEvenRecords(characterRecord []string) entity.Character {
	var character entity.Character
	readId, _ := strconv.ParseInt(characterRecord[0], 10, 64)
	if readId%2 == 0 {
		character.ID = readId
		character.Age, _ = strconv.ParseInt(characterRecord[2], 10, 64)
		character.Name = characterRecord[1]
		return character
	}
	return entity.Character{}
}

func getOddRecords(characterRecord []string) entity.Character {
	var character entity.Character
	readId, _ := strconv.ParseInt(characterRecord[0], 10, 64)
	if readId%2 != 0 {
		character.ID = readId
		character.Age, _ = strconv.ParseInt(characterRecord[2], 10, 64)
		character.Name = characterRecord[1]
		return character
	}
	return entity.Character{}
}

func (repo *characterRepository) getRecords() ([][]string, error) {
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
	return records, nil
}
