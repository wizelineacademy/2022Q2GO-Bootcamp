package repository

import (
	"context"

	models "github.com/luischitala/2022Q2GO-Bootcamp/internal/entity"
)

type Repository interface {
	ListCharacter(ctx context.Context, page uint64) ([]*models.CharacterDB, error)
	Close() error
}

func Close() error {
	return implementation.Close()
}

var implementation Repository

func SetRepository(repository Repository) {
	implementation = repository
}

func ListCharacter(ctx context.Context, page uint64) ([]*models.CharacterDB, error) {
	return implementation.ListCharacter(ctx, page)
}
