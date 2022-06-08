package repository

import (
	"encoding/csv"
	"fmt"
	"os"

	"github.com/krmirandas/2022Q2GO-Bootcamp/internal/model"
)

func ReadCSVFromUrl(url string) (data []model.Pokemon, err error) {
	csvFile, err := os.Open(url)
	if err != nil {
		return
	}

	fmt.Println("Successfully Opened CSV file")
	defer csvFile.Close()

	csvLines, err := csv.NewReader(csvFile).ReadAll()

	if err != nil {
		return
	}

	for _, row := range csvLines {
		// if idx == 0 {
		// 	continue
		// }

		newpokemon := model.Pokemon{
			ID:         row[0],
			Name:       row[1],
			Type1:      row[2],
			Type2:      row[3],
			Total:      row[4],
			HP:         row[5],
			Attack:     row[6],
			Defense:    row[7],
			SpAtk:      row[8],
			SpDef:      row[9],
			Speed:      row[10],
			Generation: row[11],
			Legendary:  row[12],
		}

		data = append(data, newpokemon)
	}
	return
}
