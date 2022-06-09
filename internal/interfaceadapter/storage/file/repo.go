package file

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"

	"github.com/jesusrevilla/capstone/internal/domain/data"
)

// Repo implements Repository Interface
type Repo struct {
	datamap map[int64]string
}

// RepoC implements RepoCofee Interface
type RepoC struct {
	coffees map[int64]data.Coffee
}

// NewRepo Constructor
func NewRepo() Repo {
	datamap := make(map[int64]string)
	return Repo{datamap}
}

// NewRepoC Constructor
func NewRepoC() RepoC {
	coffees := make(map[int64]data.Coffee)
	return RepoC{coffees}
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

// GetCoffee returs all stored coffees
func (c RepoC) GetCoffee() ([]data.Coffee, error) {
	const url = "https://random-data-api.com/api/coffee/random_coffee?size=5"

	res, err := http.DefaultClient.Get(url)
	if err != nil {
		return nil, fmt.Errorf("unable to reach [%v]: %v", url, err)
	}
	content, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("unable to read response body: %v", err)
	}

	var values []data.Coffee
	err = json.Unmarshal(content, &values)
	if err != nil {
		return nil, fmt.Errorf("unable to decode response: %v", err)
	}

	return values, nil
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
