package repository

import "github.com/2022Q2GO-Bootcamp/domain/model"

type PokemonRepository interface {
	GetAlls(u []*model.Pokemon) ([]*model.Pokemon, error)
}