package mocks

import (
	"context"
	"encoding/csv"
	"os"

	models "github.com/luischitala/2022Q2GO-Bootcamp/internal/entity"
	"github.com/stretchr/testify/mock"
)

// Repository Mock
type MockRepository struct {
	mock.Mock
}

func (mock *MockRepository) ListCharacter(ctx context.Context, page uint64) ([]*models.Character, error) {
	args := mock.Called()
	result := args.Get(0)
	return result.([]*models.Character), args.Error(1)
}

func (mock *MockRepository) Close() error {
	args := mock.Called()
	return args.Error(0)

}

type MockUseCase struct {
	mock.Mock
}

// Use Case Mock
func (mock *MockUseCase) ReadCsv() ([]*models.Character, error) {
	args := mock.Called()
	result := args.Get(0)
	return result.([]*models.Character), args.Error(1)
}

func (mock *MockUseCase) ReadCsvConcurrently(typeP string, items int, itemsPerWorker int) ([]models.Character, error) {
	args := mock.Called()
	result := args.Get(0)
	return result.([]models.Character), args.Error(1)
}

type MockCsvRepository struct {
	mock.Mock
}

func (mockCV *MockCsvRepository) ReadCsvFile() (*os.File, error) {
	args := mockCV.Called()
	result := args.Get(0)
	return result.(*os.File), args.Error(1)
}

func (mockCV *MockCsvRepository) ReadCsvFileConcurrently() ([]string, error) {
	args := mockCV.Called()
	result := args.Get(0)
	return result.([]string), args.Error(1)
}

func (mockCV *MockCsvRepository) WriteCsvFile() (*csv.Writer, *os.File, error) {
	args := mockCV.Called()
	result := args.Get(0)
	return result.(*csv.Writer), result.(*os.File), args.Error(1)
}

func (mockCV *MockCsvRepository) VerifyIfFileExists(name string) (bool, error) {
	args := mockCV.Called()
	result := args.Get(0)
	return result.(bool), args.Error(1)
}
