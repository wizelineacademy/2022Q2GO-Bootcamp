package database

import (
	"context"
	"database/sql"
	"log"
	"strconv"

	_ "github.com/lib/pq"
	models "github.com/luischitala/2022Q2GO-Bootcamp/internal/entity"
)

type PostgresRepository struct {
	db *sql.DB
}

//Constructor

func NewPostgresRepository(url string) (*PostgresRepository, error) {
	db, err := sql.Open("postgres", url)
	if err != nil {
		return nil, err
	}
	return &PostgresRepository{db}, nil
}

func (repo *PostgresRepository) ListCharacter(ctx context.Context, id uint64) ([]*models.Character, error) {
	query := ""
	parsedInt := ""
	if id != 0 {
		parsedInt = strconv.FormatUint(id, 10)
		query = "SELECT id, name, species, type, gender, image, url, created FROM public.characters where id = " + parsedInt
	} else {
		query = "SELECT id, name, species, type, gender, image, url, created FROM public.characters "
	}
	rows, err := repo.db.QueryContext(ctx, query)

	if err != nil {
		return nil, err
	}
	defer func() {
		err = rows.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	var characters []*models.Character
	for rows.Next() {
		var character = models.Character{}
		if err = rows.Scan(&character.Id, &character.Name, &character.Species, &character.Type, &character.Gender, &character.Image, &character.Url, &character.Created); err == nil {
			characters = append(characters, &character)

		}

	}

	if err = rows.Err(); err != nil {

		return nil, err
	}

	return characters, nil
}

func (repo *PostgresRepository) Close() error {
	return repo.db.Close()
}
