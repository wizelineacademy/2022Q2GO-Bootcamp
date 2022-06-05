package interactor

import (
	"github.com/2022Q2GO-Bootcamp/domain/model"
	"github.com/2022Q2GO-Bootcamp/usecase/presenter"
	"github.com/2022Q2GO-Bootcamp/usecase/repository"
)

type pokemonInteractor struct {
	PokemonRepository repository.PokemonRepository
	PokemonPresenter  presenter.PokemonPresenter
}

type UserInteractor interface {
	Get(u []*model.Pokemon) ([]*model.Pokemon, error)
}

func NewPokemonInteractor(r repository.PokemonRepository, p presenter.PokemonPresenter) UserInteractor {
	return &pokemonInteractor{r, p}
}

func (us *pokemonInteractor) Get(u []*model.Pokemon) ([]*model.Pokemon, error) {
	u, err := us.PokemonRepository.GetAlls(u)
	if err != nil {
		return nil, err
	}

	return us.PokemonPresenter.ResponsePokemons(u), nil
}