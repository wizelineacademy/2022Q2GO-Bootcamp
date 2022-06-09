package controller

import (
	"context"
	"encoding/csv"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"strings"
	"testing"
	"time"

	mocks "github.com/luischitala/2022Q2GO-Bootcamp/internal/controller/mocks"
	models "github.com/luischitala/2022Q2GO-Bootcamp/internal/entity"
	"github.com/luischitala/2022Q2GO-Bootcamp/internal/infrastructure/database"
	"github.com/luischitala/2022Q2GO-Bootcamp/internal/usecase"
	"github.com/stretchr/testify/assert"
)

func TestCharacter(t *testing.T) {
	file := "../../files/characters.csv"
	//io util opened file
	csvFileIo, err := ioutil.ReadFile(file)
	if err != nil {
		log.Printf("Error while opening the CSV file: %s", err)
	}
	//Os opened file
	csvFile, err := os.Open(file)
	if err != nil {
		log.Printf("Error while opening the CSV file: %s", err)
	}
	assert.Nil(t, err)
	// Return the content of the csv file
	content := strings.Split(string(csvFileIo), "\n")
	defer csvFile.Close()
	// Seed to try even or odd
	rand.Seed(time.Now().Unix())
	typeO := []string{
		"even",
		"odd",
	}
	n := rand.Int() % len(typeO)
	option := typeO[n]
	isCorrect := true
	testTable := []struct {
		name          string
		exampleModel  *models.Character
		expectedArray []models.Character
		function      string
		expectedErr   error
		typeParam     string
		databaseConn  string
	}{
		{
			name:          "database-connection-characters",
			function:      "Characters",
			exampleModel:  &mocks.Character,
			expectedArray: nil,
			expectedErr:   nil,
			typeParam:     "",
			databaseConn:  "postgres://postgres:postgres@db:5432/postgres?sslmode=disable",
		},
		{
			name:          "test-read-csv",
			function:      "ReadCsv",
			exampleModel:  &mocks.Character,
			expectedArray: nil,
			expectedErr:   nil,
			typeParam:     "",
			databaseConn:  "",
		},
		{
			name:          "test-read-csv-concurrently",
			function:      "ReadCsvConcurrently",
			exampleModel:  &mocks.Character,
			expectedArray: []models.Character{},
			expectedErr:   mocks.ErrBadStatusCode,
			typeParam:     option,
			databaseConn:  "",
		},
		{
			name:          "write-on-csv",
			function:      "WriteCsv",
			exampleModel:  &mocks.Character,
			expectedArray: []models.Character{},
			expectedErr:   nil,
			typeParam:     option,
			databaseConn:  "",
		},
	}

	for _, tc := range testTable {
		t.Run(tc.name, func(t *testing.T) {
			mockCsvRepo := new(mocks.MockCsvRepository)
			characterUseCase, _ := usecase.NewCharacterUseCase(mockCsvRepo)
			// Method Execution
			switch tc.function {
			case "characters":
				postgresRepository, err := database.NewPostgresRepository(tc.databaseConn)
				if err != nil {
					log.Println(err)
				}
				assert.Nil(t, err)

				listCharacters, err := postgresRepository.ListCharacter(context.Background(), 2)
				if err != nil {
					log.Println(err)
				}
				assert.NotNil(t, listCharacters)
			case "ReadCsv":
				mockCsvRepo.On("ReadCsvFile").Return(csvFile, nil)
				result, err := characterUseCase.ReadCsv()
				if err != nil {
					log.Printf("Error during the character transformation from csv: %s", err)
				}
				assert.Nil(t, err)
				firstRow := result[0]
				assert.Equal(t, tc.exampleModel, firstRow)
			case "ReadCsvConcurrently":

				mockCsvRepo.On("ReadCsvFileConcurrently").Return(content, nil)
				result, err := characterUseCase.ReadCsvConcurrently(tc.typeParam, 9, 3)
				if err != nil {
					log.Printf("Error during the character transformation from csv: %s", err)
				}
				assert.Nil(t, err)
				assert.NotEmpty(t, result, tc.expectedArray)
				// Verify if the result
				switch tc.typeParam {

				case "odd":
					if result[1].Id%2 == 0 {
						isCorrect = false
						log.Println("Type query parameter is having some problems with odd")

					}
				case "even":
					if result[1].Id%2 != 0 {
						isCorrect = false
						log.Println("Type query parameter is having some problems with even")

					}
				}
				assert.True(t, isCorrect)
			case "WriteCsv":
				createdFile, _ := os.Create("../../files/characterResult.csv")
				csvwriter := csv.NewWriter(createdFile)

				mockCsvRepo.On("VerifyIfFileExists").Return(true, nil)
				mockCsvRepo.On("WriteCsvFile").Return(*csvwriter, *createdFile, nil)
				// apiResponse, err := characterUseCase.ConsultApi()
				apiResponse, err := characterUseCase.ConsultApi()
				if err != nil {
					log.Printf("Error during the character transformation from csv: %s", err)
				}
				assert.NotEmpty(t, apiResponse)

			}

		})
	}

}
