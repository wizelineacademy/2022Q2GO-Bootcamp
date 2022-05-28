package usecase

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"

	models "github.com/luischitala/2022Q2GO-Bootcamp/internal/entity"
	"github.com/luischitala/2022Q2GO-Bootcamp/internal/repository"
)

//Interface that allows to execute all the entity operations
type Character interface {
	Reader
	Writer
}

type Reader interface {
	ReadCsv() ([]models.Character, error)
}

type Writer interface {
	WriteCsv() (models.CharacterResponse, error)
}

//Chain struct to separate logic between the next layer
type rs struct {
	Csv repository.Csv
}

func NewCharacterUseCase(rcsv repository.Csv) Character {
	return &rs{
		rcsv,
	}
}

func (r *rs) ReadCsv() ([]models.Character, error) {
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
			Id:      id,
			Name:    line[1],
			Status:  line[2],
			Species: line[3],
			Type:    line[4],
			Gender:  line[5],
			Image:   line[6],
			Url:     line[7],
			Created: line[8],
		}
		characters = append(characters, character)

	}
	return characters, nil
}

func (r *rs) WriteCsv() (models.CharacterResponse, error) {
	// Client
	cli := http.Client{Timeout: time.Duration(1) * time.Second}
	req, err := http.NewRequest("GET", "https://rickandmortyapi.com/api/character", nil)
	if err != nil {
		fmt.Printf("error %s", err)
	}
	req.Header.Add("Accept", `application/json`)
	// Request exec
	resp, err := cli.Do(req)

	if err != nil {
		log.Fatalln(err)
	}
	// Close response
	defer resp.Body.Close()

	bodyBytes, _ := ioutil.ReadAll(resp.Body)

	// Convert response body to string
	bodyString := string(bodyBytes)
	fmt.Println("API Response as String:\n" + bodyString)

	csvWriter, csvFile, err := r.Csv.WriteCsvFile()
	// Convert response body to CaracterResponse struct
	var charStruct models.CharacterResponse
	json.Unmarshal(bodyBytes, &charStruct)
	fmt.Printf("API Response as struct %+v\n", charStruct)
	// Close and flush the csv
	defer csvFile.Close()
	defer csvWriter.Flush()
	// Iterate the response
	for _, character := range charStruct.Results {
		// Populate the row
		row := []string{
			strconv.Itoa(character.Id),
			character.Name,
			character.Status,
			character.Species,
			character.Type,
			character.Gender,
			character.Image,
			character.Url,
			character.Created,
		}
		// Write the row
		if err := csvWriter.Write(row); err != nil {
			log.Fatalln("error writing record to file", err)
		}
	}

	return charStruct, nil
}
