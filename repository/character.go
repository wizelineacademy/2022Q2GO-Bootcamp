package repository

import (
	"context"
	"log"

	"github.com/TanZng/toh-api/models"
)

type CharacterRepository interface {
	InsertCharacter(ctx context.Context, user *models.Character) error
	GetCharacterById(ctx context.Context, id int64) (*models.Character, error)
	Close() error
}

var implementation CharacterRepository

func SetRepository(repository CharacterRepository) {
	implementation = repository
}

func InsertCharacter(ctx context.Context, user *models.Character) error {
	return implementation.InsertCharacter(ctx, user)
}

func GetCharacterById(ctx context.Context, id int64) (*models.Character, error) {

	c, err := implementation.GetCharacterById(ctx, id)
	if err != nil {
		log.Fatal("ERROR: GetCharacterById", err)
		return nil, err
	}
	// log.Println("Character:", c)
	return c, nil
}

func Close() error {
	return implementation.Close()
}
