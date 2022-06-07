package repository
import (
    "log"
	"fmt"
	"os" 
	"encoding/csv"
	"strconv"
	"github.com/waltergomezda/2022Q2GO-Bootcamp/internal/entity"
)

type  CsvRepository interface {
	ReadRecords()([]entity.CsvRecord, error)
	WriteRecords(records []entity.CatFact)(error)
}
//Define struct for repo
type csvRepository struct {
	filePath string
}
//Constructor
func NewCsvRepository(file string) CsvRepository{
	return &csvRepository{filePath:file}
}
// Interface implementation
func (r *csvRepository) ReadRecords() ([]entity.CsvRecord, error){
	// open file
    f, err := os.Open(r.filePath)
    if err != nil {
        log.Fatal(err)
		return nil, err
    }

    // remember to close the file at the end of the program
    defer f.Close()

    // read csv values using csv.Reader
    csvReader := csv.NewReader(f)
    data, err := csvReader.ReadAll()
    if err != nil {
        log.Fatal(err)
		return nil, err
    }
	
    // convert csv records to array of structs
    results := createCsvRecordList(data)
	fmt.Println(results)
	return results, err
}

func (r *csvRepository) WriteRecords(records []entity.CatFact)(err error){
	
	csvFile, err := os.Create(r.filePath)
	if err != nil {
		log.Fatal("failed creating file: %s", err)
	}
	csvwriter := csv.NewWriter(csvFile)
	for _, row := range records{
		fmt.Println([]string{row.Fact})
		_ = csvwriter.Write([]string{row.Fact})
	} 
	csvwriter.Flush()
	csvFile.Close()
	return err
}

func createCsvRecordList(data [][]string) ([]entity.CsvRecord) {
    var csvRecordList []entity.CsvRecord
    for _, line := range data {
		var rec entity.CsvRecord
		
		for j, field := range line {
			if j == 0 {
				key, _ := strconv.ParseInt(field, 10, 64)
				rec.Key = key
			} else if j == 1 {
				rec.Value = field
			}
		}
		if rec.Key > 0 { // Omit items with invalid key
			csvRecordList = append(csvRecordList, rec)
		}
    }
    return csvRecordList
}