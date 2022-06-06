package csv

import (
	"encoding/csv"
	"fmt"
	"github.com/esvarez/go-api/internal/boardgame"
	"github.com/esvarez/go-api/internal/pokemon"
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
	// TODO add index dynamically
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

func getCSVWriter(filePath string) *csv.Writer {
	f, err := os.OpenFile(filePath, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		log.Fatalf("failed opening file: %s", err)
	}

	return csv.NewWriter(f)
}

func (c Client) FindBoardGame(id int) (*boardgame.BoardGame, error) {
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

func (c Client) GetAllPokemon() ([]pokemon.Pokemon, error) {
	csvClient := getCSVReader(c.path)

	data, err := csvClient.ReadAll()
	if err != nil {
		return nil, err
	}
	pokemons := []pokemon.Pokemon{}
	for _, record := range data[1:] {
		id, err := strconv.Atoi(record[0])
		if err != nil {
			return nil, err
		}
		pokemons = append(pokemons, pokemon.Pokemon{
			ID:   id,
			Name: record[1],
		})
	}
	return pokemons, nil
}

func (c Client) AddPokemon(pokemon *pokemon.Pokemon) error {
	writer := getCSVWriter(c.path)
	/*
		reader := getCSVReader(c.path)

		records, err := reader.ReadAll()
		//TODO improve error handling
		if err != nil {
			return err
		}
	*/
	// records = append(records, []string{strconv.Itoa(pokemon.ID), pokemon.Name})
	// TODO avoid duplicates
	if err := writer.Write([]string{strconv.Itoa(pokemon.ID), pokemon.Name}); err != nil {
		return err
	}
	writer.Flush()

	return nil
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
