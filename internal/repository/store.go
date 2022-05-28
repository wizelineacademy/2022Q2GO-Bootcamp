package repository

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"time"
)

func CreateCSV(empData map[string]interface{}) {
	now := time.Now() // current local time
	sec := now.Unix() // number of seconds since January 1, 1970 UTC
	csvFile, err := os.Create(fmt.Sprintf("exported_%d.csv", sec))

	if err != nil {
		log.Fatalf("failed creating file: %s", err)
	}

	csvwriter := csv.NewWriter(csvFile)

	// jsonByte, _ := json.Marshal(empData)
	// jsonString := string(jsonByte)
	// for _, empRow := range rates {
	// 	// empRow, empRow := empData[key].([]string)
	// 	fmt.Println((empRow))
	// 	_ = csvwriter.Write(empRow)
	// }

	csvwriter.Flush()
	csvFile.Close()
}
