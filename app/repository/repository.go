package repository

import (
	"context"

	"github.com/luischitala/2022Q2GO-Bootcamp/models"
)

type Repository interface {
	ListCharacter(ctx context.Context, page uint64) ([]*models.Character, error)
	Close() error
}

func Close() error {
	return implementation.Close()
}

var implementation Repository

func SetRepository(repository Repository) {
	implementation = repository
}

func ListCharacter(ctx context.Context, page uint64) ([]*models.Character, error) {
	return implementation.ListCharacter(ctx, page)
}
