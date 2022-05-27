package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Fish struct {
	Id   int64  `json:"id" uri:"id"`
	Name string `json:"name"`
}

type MyFish struct {
	Items []Fish
}

//External API URL
const (
	url = "https://quiet-atoll-32754.herokuapp.com/fish"
)

// Http Client to get Fish data from External API
func runClient(c *gin.Context) {
	resp, getErr := http.Get(url)
	if getErr != nil {
		log.Fatal(getErr)
	}
	body, readErr := ioutil.ReadAll(resp.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}

	apiFishData := MyFish{}
	if err := json.Unmarshal(body, &apiFishData.Items); err != nil {
		fmt.Printf("error unmarshaling JSON: %v\n", err)
	}
	//Save data to CSV file
	csvWriter(apiFishData)
	//JSON Response
	c.IndentedJSON(http.StatusOK, apiFishData.Items)
}

// getFish responds with the list of all fish as JSON.
func getFish(c *gin.Context) {
	fish := csvReader()
	c.IndentedJSON(http.StatusOK, fish.Items)
}

// getFishById responds with the fish by Id as JSON.
func getFishById(c *gin.Context) {
	myFish := csvReader()
	idParam, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		fmt.Println(err)
	}
	var fish Fish
	for _, fishItem := range myFish.Items {
		if fishItem.Id == idParam {
			fish = fishItem
		}
	}
	if (Fish{}) == fish {
		c.JSON(http.StatusNotFound, "Resource Not Found: 404")
	} else {
		c.IndentedJSON(http.StatusOK, fish)
	}
}

func csvReader() MyFish {
	// 1. Open the file
	recordFile, err := os.Open("./fish.csv")
	if err != nil {
		fmt.Println("CSV not valid: ", err)
	}
	// 2. Initialize the reader
	reader := csv.NewReader(recordFile)
	// 3. Read all the records
	records, err := reader.ReadAll()
	if err != nil {
		fmt.Println("An error ocurred while reading the file: ", err)
	}
	// 4. Add records to MyFish struct
	data := MyFish{}
	for _, row := range records {
		id, err := strconv.ParseInt(row[0], 10, 64)
		if err != nil {
			fmt.Println(err)
		}
		item := Fish{Id: id, Name: row[1]}
		data.Items = append(data.Items, item)
	}
	//5. Close csv file
	err = recordFile.Close()
	if err != nil {
		fmt.Println("An error encountered closing the csv file ", err)
	}
	return data
}

func csvWriter(myFish MyFish) {
	// 1. Open the file
	recordFile, err := os.Create("./fishFromAPI.csv")
	if err != nil {
		fmt.Println("An error encountered:", err)
	}

	// 2. Initialize the writer
	writer := csv.NewWriter(recordFile)

	// 3. Write all the records from myFish
	var data [][]string
	for _, record := range myFish.Items {
		row := []string{strconv.FormatInt(record.Id, 10), record.Name}
		data = append(data, row)
	}
	writer.WriteAll(data)

	recordFile.Close()
}

func main() {
	router := gin.Default()
	router.GET("/fish", getFish)
	router.GET("/fish/:id", getFishById)
	router.GET("/external-api", runClient)

	router.Run("localhost:8080")
}
