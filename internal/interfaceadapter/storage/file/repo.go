package file

import (
	"encoding/csv"
	"os"
	"strconv"

	"github.com/jesusrevilla/capstone/internal/domain/data"
)

// Repo implements Repository Interface
type Repo struct {
	datamap map[int64]string
}

// NewRepo Constructor
func NewRepo() Repo {
	datamap := make(map[int64]string)
	return Repo{datamap}
}

//Find Returns data with the provided id
func (r Repo) Find(id int64) (*data.Data, error) {
	err := fillDatamap(r.datamap)
	if err != nil {
		return &data.Data{Id: 0, Item: ""}, err
	}
	data := data.Data{Id: id, Item: r.datamap[id]}
	return &data, nil
}

func fillDatamap(datamap map[int64]string) error {
	f, err := os.Open("../../internal/interfaceadapter/storage/file/data.csv")
	if err != nil {
		return err
	}

	defer f.Close()

	csvReader := csv.NewReader(f)
	records, err := csvReader.ReadAll()
	if err != nil {
		return err
	}
	for _, value := range records {
		key, err := strconv.ParseInt(value[0], 10, 64)
		if err != nil {
			return err
		}
		datamap[key] = value[1]
	}
	return nil
}
