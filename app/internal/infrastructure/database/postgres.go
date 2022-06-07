package database

import (
	"context"
	"database/sql"
	"fmt"
	"log"

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

func (repo *PostgresRepository) ListCharacter(ctx context.Context, page uint64) ([]*models.CharacterDB, error) {
	//Only allow edit for owners
	fmt.Println("Hola")

	rows, err := repo.db.QueryContext(ctx, "SELECT * FROM characters LIMIT $1 OFFSET $2", 2, page*2)
	if err != nil {
		return nil, err
	}
	defer func() {
		err = rows.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	var characters []*models.CharacterDB
	for rows.Next() {
		var character = models.CharacterDB{}
		if err = rows.Scan(&character.Id, &character.Name); err == nil {
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
