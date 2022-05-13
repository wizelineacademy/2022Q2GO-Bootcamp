package csv

import (
	"encoding/csv"
	"fmt"
	"github.com/esvarez/go-api/internal/boardgame"
	errs "github.com/esvarez/go-api/pkg/error"
	"log"
	"os"
	"strconv"
)

const (
	idI = iota
	nameI
	descriptionI
	minPlayersI
	maxPlayersI
	DurationI
)

type Client struct {
	path string
}

func NewCSVClient(filePath string) *Client {
	getCSVReader(filePath)
	return &Client{
		path: filePath,
	}
}

func getCSVReader(filePath string) *csv.Reader {
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatalf("failed opening file: %s", err)
	}

	return csv.NewReader(f)
}

func (c *Client) FindBoardGame(id int) (*boardgame.BoardGame, error) {
	csvClient := getCSVReader(c.path)

	data, err := csvClient.ReadAll()
	if err != nil {
		return nil, err
	}

	record, err := getBoardGame(id, data)
	if err != nil {
		return nil, err
	}

	return &boardgame.BoardGame{
		ID:          record[idI],
		Name:        record[nameI],
		Description: record[descriptionI],
		MinPlayers:  record[minPlayersI],
		MaxPlayers:  record[maxPlayersI],
		Duration:    record[DurationI],
	}, nil
}

func getBoardGame(id int, data [][]string) ([]string, error) {
	d := data
	for _, record := range d[1:] {
		idbg, err := strconv.Atoi(record[idI])
		if err != nil {
			// TODO log error for invalid entries
			continue
		}
		if idbg == id {
			return record, nil
		}
	}
	return nil, fmt.Errorf("boardgame not found : %w", errs.ErrNotFound)
}
