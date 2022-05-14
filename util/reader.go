package util

import (
	"encoding/csv"
	"net/http"
	errorhandler "./errorhandler"
)

func ReadCSVFromUrl(url string) ([][]string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, errorhandler.ErrInternalServerError
	}
	
	defer resp.Body.Close()
	reader := csv.NewReader(resp.Body)
	reader.Comma = ','
	data, err := reader.ReadAll()
	if err != nil {
		return nil, errorhandler.ErrInternalServerError
	}

	return data, nil
}