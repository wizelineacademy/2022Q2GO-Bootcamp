package repository

import (
	"encoding/csv"
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/krmirandas/2022Q2GO-Bootcamp/internal/hook/operationhandler"
	"github.com/krmirandas/2022Q2GO-Bootcamp/internal/model"
)

func CreateCSV(empData model.ZipCodeInfo) error {
	var csvFile *os.File
	var err error
	fileName := "./data/exported.csv"

	if _, err := os.Stat(fileName); err == nil {
		// path/to/whatever exists
		csvFile, err = os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY, 0644)
	} else if errors.Is(err, os.ErrNotExist) {
		// path/to/whatever does *not* exist
		csvFile, err = os.Create(fmt.Sprintf(fileName))
	}

	if err != nil {
		log.Fatalf("failed creating file: %s", err)
		return err
	}

	m := make(map[string]model.ZipCodeInfo)
	r := empData
	m[empData.PostCode] = r

	csvwriter := csv.NewWriter(csvFile)

	for _, v := range m {
		r := make([]string, 0, 4)
		r = append(r, v.PostCode)
		r = append(r, v.Country)
		r = append(r, v.CountryAbbreviation)
		if v.Places != nil {
			for _, v2 := range v.Places {
				for _, v3 := range v2 {
					r = append(r, fmt.Sprintf("%v", v3))
				}
			}
		}

		log.Println(*operationhandler.SuccessfullyAdded)
		err := csvwriter.Write(r)
		if err != nil {
			return err
		}
	}

	csvwriter.Flush()
	defer csvFile.Close()
	return nil
}
