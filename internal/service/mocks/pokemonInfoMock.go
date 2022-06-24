package mocks

import (
	"github.com/krmirandas/2022Q2GO-Bootcamp/internal/entity"
	"github.com/stretchr/testify/mock"
)

type mockPokemonInfoRepo struct {
	mock.Mock
}

func (mr *mockPokemonInfoRepo) WritePokemon(pokemonInfo entity.PokemonInfo) error {
	arg := mr.Called(pokemonInfo)
	return arg.Error(0)
}

func (mr *mockPokemonInfoRepo) ConcuRSwWP(itemsPerWorker int) []entity.PokemonInfo {
	arg := mr.Called(itemsPerWorker)
	return arg.Get(0).([]entity.PokemonInfo)
}

type NewPokemonInfoRepositoryT interface {
	mock.TestingT
}

func NewPokemonInfoRepository(t NewPokemonInfoRepositoryT) *mockPokemonInfoRepo {
	mock := &mockPokemonInfoRepo{}
	mock.Mock.Test(t)

	mock.AssertExpectations(t)

	return mock
}
