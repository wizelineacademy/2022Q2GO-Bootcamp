package mocks

import (
	"github.com/krmirandas/2022Q2GO-Bootcamp/internal/entity"
	"github.com/stretchr/testify/mock"
)

type mockPokemonRepo struct {
	mock.Mock
}

func (mr *mockPokemonRepo) Count() (int, error) {
	arg := mr.Called()
	return arg.Int(0), arg.Error(1)
}

func (mr *mockPokemonRepo) ReadOnePokemon(id string) (entity.Pokemon, error) {
	arg := mr.Called(id)
	return arg.Get(0).(entity.Pokemon), arg.Error(1)
}

func (mr *mockPokemonRepo) ReadPokemon() ([]entity.Pokemon, error) {
	arg := mr.Called()
	return arg.Get(0).([]entity.Pokemon), arg.Error(1)
}

type NewPokemonRepositoryT interface {
	mock.TestingT
}

func NewPokemonRepository(t NewPokemonRepositoryT) *mockPokemonRepo {
	mock := &mockPokemonRepo{}
	mock.Mock.Test(t)

	mock.AssertExpectations(t)

	return mock
}
