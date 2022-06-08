package usecase

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	models "github.com/luischitala/2022Q2GO-Bootcamp/internal/entity"
	"github.com/luischitala/2022Q2GO-Bootcamp/internal/repository"
)

//Interface that allows to execute all the entity operations
type Character interface {
	Reader
	ReaderConcurrently
	Writer
	ApiClient
}

type Reader interface {
	ReadCsv() ([]*models.Character, error)
}
type ReaderConcurrently interface {
	ReadCsvConcurrently(string, int, int) ([]models.Character, error)
}

type Writer interface {
	WriteCsv([]byte) (models.CharacterResponse, error, bool)
}

type ApiClient interface {
	ConsultApi() ([]byte, error)
}

//Chain struct to separate logic between the next layer
type rs struct {
	Csv repository.Csv
}

func NewCharacterUseCase(rcsv repository.Csv) (Character, error) {
	return &rs{
		rcsv,
	}, nil
}

// Function that contains the rules to iterate a csv file from the repository, returns a list of characters
func (r *rs) ReadCsv() ([]*models.Character, error) {
	characters := make([]*models.Character, 0)
	// Open the csv file
	csvFile, err := r.Csv.ReadCsvFile()
	if err != nil {
		log.Println(err)
		return nil, err
	}
	// Create the csv reader
	csvLines, err := csv.NewReader(csvFile).ReadAll()
	// Handler the csv reader errors
	if err != nil {
		log.Println(err)
		return nil, err
	}
	// Iterate the csv rows to populate a character struct
	for _, line := range csvLines {
		id, _ := strconv.Atoi(line[0])

		character := &models.Character{
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

// Function that send a get request to external api and returns the response
func (r *rs) ConsultApi() ([]byte, error) {
	// Client
	cli := http.Client{Timeout: time.Duration(1) * time.Second}
	req, err := http.NewRequest("GET", "https://rickandmortyapi.com/api/character", nil)
	// Handle if the request construction is wrong
	if err != nil {
		log.Printf("error %s", err)
		return nil, err
	}
	// Request exec
	req.Header.Add("Accept", `application/json`)
	resp, err := cli.Do(req)
	// Handle if the request has an error
	if err != nil {
		log.Println(err)
		return nil, err
	}
	// Close response
	defer resp.Body.Close()

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	// Handle if the body response has errors
	if err != nil {
		log.Println(err)
		return nil, err

	}
	// Log the response
	log.Println("API Response as String:\n" + string(bodyBytes))

	return bodyBytes, nil
}

// Function that contains the rules to consult and external api, to write the result into a csv file
func (r *rs) WriteCsv(apiResponse []byte) (models.CharacterResponse, error, bool) {
	// Initialize the response struct
	var charStruct models.CharacterResponse

	// Verify if the csv file exists
	exists, err := r.Csv.VerifyIfFileExists("files/characterResult.csv")
	if !exists {
		log.Println(err)

	}
	// Generate the csv writer
	csvWriter, csvFile, err := r.Csv.WriteCsvFile()
	if err != nil {
		log.Println(err)
		return charStruct, nil, exists
	}
	// Convert response body to CaracterResponse struct
	json.Unmarshal(apiResponse, &charStruct)
	log.Printf("API Response as struct %+v\n", charStruct)
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

	return charStruct, nil, exists
}

// Function that contains the rules to iterate a csv file from the repository, returns a list of characters concurrently
func (r *rs) ReadCsvConcurrently(typeP string, items int, itemsPerWorker int) ([]models.Character, error) {
	// Open CSV file
	content, err := r.Csv.ReadCsvFileConcurrently()
	if err != nil {
		return nil, err
	}
	// Handle the error if the csv file is empty
	if len(content) == 0 {
		return nil, fmt.Errorf("the character's file is empty")
	}
	//Variables, response and channel objects
	itemsF := float64(itemsPerWorker)
	characters := []models.Character{}
	itemsI := items
	workers := 1.0
	i := 0
	realCounter := 0
	realItems := 0
	diff := 0
	var workChan []chan string

	// Check if the items and items per worker are greater than 0
	if items > 0 && itemsPerWorker > 0 {
		// Get the real amount of items to consult based on the type parameter
		// This evaluation will provide how many records will be necessary for the workers to read by multiples of the items_per_worker
		if typeP == "even" {
			if items%2 == 0 {
				for itemsI > 0 {
					realCounter += 1
					itemsI -= 2
				}
				diff = items - realCounter
				realItems = items + (diff * 2)
			}
		}
		if typeP == "odd" {
			if items%2 != 0 {
				for itemsI > 0 {
					if itemsI != 1 {
						itemsI -= 2
						realCounter += 1
					} else {
						itemsI = 0
						realCounter += 1
					}
				}
				diff = items - realCounter
				realItems = items + (diff * 2) + 1
			}
		}

		// Get the quantity of workers if the items from the query parameter is less than the real content of the csv file
		if items < len(content) {

			workers = float64(realItems) / float64(itemsPerWorker)

			// Else, get the quantity of workers based on the total items from the csv file
		} else {
			workers = float64(len(content)) / float64(itemsPerWorker)
		}
		// Asign the work channel with 0
		workChan = make([]chan string, 0)
		// Populate the work channel for the even numbers
		for ij := 0; ij < int(workers); ij++ {
			realChan := make(chan string, itemsPerWorker)
			workChan = append(workChan, realChan)
		}

		//Recalculate the final items for an odd number
		itemsF = (workers - float64(int(workers))) * float64(itemsPerWorker)
		if itemsF > 0 {
			realChan := make(chan string, int(itemsF))
			workChan = append(workChan, realChan)
		}
		// If one of the parameters is 0 asign the real chan with 0
	} else {
		workChan = make([]chan string, 0)
		realChan := make(chan string)
		workChan = append(workChan, realChan)
	}

	i = 0
	// Range over the work channel
	for worker, row := range workChan {
		// Send final workload, last items or item
		if worker == len(workChan)-1 && itemsF > 0 {
			ItemsR := content[i : i+int(itemsF)]
			go appendWorkload(row, typeP, ItemsR)
			log.Printf("Worker %d finished. Read %d items \n", worker, len(ItemsR))
			i = i + int(itemsF)

			// Send workload for the multiples of items_per_worker
		} else {
			ItemsR := content[i : i+itemsPerWorker]
			go appendWorkload(row, typeP, ItemsR)
			log.Printf("Worker %d finished. Read %d items \n", worker, len(ItemsR))
			i = i + itemsPerWorker
		}
	}
	// Iterate the workload sent to the working channel
	for _, j := range workChan {
		// Iterate the rows from the workload
		for row := range j {

			// Split the received string
			line := strings.Split(row, ",")

			// Parse the id
			id, err := strconv.Atoi(line[0])
			// Handle the 0
			if id == 0 {
				id = 1

			}
			// Handle the string to int conversion, if is not the error from parsing 0
			if err != nil && id != 1 {
				log.Printf("Error converting the id %s", line[0])
				break
			}
			// Call the retrieveCharacter function that will split and read the string to populate the character and then return it
			char, err := retrieveCharacter(id, row)
			// Handle if the character is not found or if there was an error
			if err != nil {
				log.Printf("Error retrieving the character with id %d ", id)
				break
			}
			// Populate the character array
			characters = append(characters, char)
		}
	}
	return characters, nil
}

// Function that's executed as a go routine to distribute the workload between the workers if the query parameter is odd or even, sent to the worker's channel
func appendWorkload(row chan string, typeP string, af []string) {

	for _, l := range af {
		// Separate the array of characters
		line := strings.Split(l, ",")

		// Parsing the id comming from the first column of the csv file rows
		id, err := strconv.Atoi(line[0])
		if id == 0 {
			id = 1
		}
		if err != nil && id != 1 {
			log.Printf("Error converting the id %s", line[0])
			break
		}
		switch typeP {
		case "even":
			if id%2 == 0 {
				row <- l
			}
		case "odd":
			if id%2 != 0 {
				row <- l

			}

		default:
			row <- l

		}
	}

	close(row)
}

// Function to receive the string sent to the channel to split the csv row and populate and return the character struct
func retrieveCharacter(id int, row string) (models.Character, error) {
	// Initialize the Character
	a := models.Character{}
	// Split the sent string
	line := strings.Split(row, ",")
	// Populate the character model with the row information
	a = models.Character{
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
	// Return the character
	return a, nil
}
