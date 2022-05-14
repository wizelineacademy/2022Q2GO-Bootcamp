package presenter

import "github.com/2022Q2GO-Bootcamp/domain/model"

type PokemonPresenter interface {
	ResponsePokemons(u []*model.Pokemon) []*model.Pokemon
}
