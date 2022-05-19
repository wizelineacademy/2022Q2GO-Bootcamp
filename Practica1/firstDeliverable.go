package main

import (
	"encoding/csv"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
)

type data struct {
	ID   int
	Name string
}

func main() {
	http.HandleFunc("/readCSVFile", handlerReadFile)
	http.HandleFunc("/searchID", handlerSearchID)
	fmt.Println("The server is running")

	/*e := http.ListenAndServe(":8090", nil)

	if e != nil {
		log.Fatal(e)
	}*/
	go http.ListenAndServe(":8090", nil)

	res, err := http.Get("http://127.0.0.1:8090/readCSVFile")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("res.StatusCode: %d\n", res.StatusCode)

	if res.StatusCode != 200 {
		b, _ := ioutil.ReadAll(res.Body)
		log.Fatal(string(b))
	}

	fmt.Println("")
	fmt.Println("")

	res2, err := http.Get("http://127.0.0.1:8090/searchID")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("res2.StatusCode: %d\n", res2.StatusCode)

	if res2.StatusCode != 200 {
		b, _ := ioutil.ReadAll(res.Body)
		log.Fatal(string(b))
	}

}

func handlerReadFile(w http.ResponseWriter, r *http.Request) {
	file := "./informacion.csv"
	csvLines, e := readData(file)

	if e != nil {
		fmt.Print("Error trying to get the info in CSV file ")
		log.Fatal(e)
	}
	fmt.Println("Successfully got the info in CSV file")

	var dataInfo []data
	for _, line := range csvLines {
		idInt, e := strconv.Atoi(line[0])

		if e != nil {
			fmt.Printf("%T \n %v", idInt, idInt)
			//return e
			log.Fatal(e)
		}

		d := data{
			ID:   idInt,
			Name: line[1],
		}

		dataInfo = append(dataInfo, d)

	}
	fmt.Println("Information: ", dataInfo)

	dataLines := fmt.Sprint(csvLines)

	information := []byte(dataLines)
	_, err := w.Write(information)

	if err != nil {
		log.Fatal()
	}
	fmt.Println("Successfully completed the task")
}

func handlerSearchID(w http.ResponseWriter, r *http.Request) {
	file := "./informacion.csv"
	IDToSearch := 2
	csvLines, e := readData(file)

	if e != nil {
		fmt.Print("Error trying to get the info in CSV file ")
		log.Fatal(e)
	}
	fmt.Println("Successfully got the info in CSV file")

	var dataID []data
	for _, line := range csvLines {
		idInt, e := strconv.Atoi(line[0])

		if e != nil {
			fmt.Printf("%T \n %v", idInt, idInt)
			//return e
			log.Fatal(e)
		}

		d := data{
			ID:   idInt,
			Name: line[1],
		}
		if idInt == IDToSearch {
			dataID = append(dataID, d)
			fmt.Println(dataID)
		}

	}
	fmt.Println("Datos: ", dataID)

	dataLines := fmt.Sprint(dataID)

	information := []byte(dataLines)
	_, err := w.Write(information)

	if err != nil {
		log.Fatal()
	}
	fmt.Println("Successfully completed the task")
}

func readData(fileName string) ([][]string, error) {
	f, err := os.Open(fileName) // open file

	if err != nil {
		return [][]string{}, err
	}

	fmt.Println("Successfully opened the file")
	// close the file at the end of the program
	defer f.Close()

	// read csv values using csv.Reader
	csvReader := csv.NewReader(f)

	// skip first line
	if _, err := csvReader.Read(); err != nil {
		return [][]string{}, err
	}

	//Here we skip the first line, which contains the column names.
	csvLines, err := csvReader.ReadAll()

	if err != nil {
		log.Fatal(err)
	}
	return csvLines, nil
}
