package service

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"log"
	"testing"
	"toh-api/internal/entity"
	"toh-api/internal/repository"
	"toh-api/internal/service/mocks"
	"toh-api/test/testdata"
)

// mockCharacterRepository custom/manual mocked repository
type mockCharacterRepository struct {
	mock.Mock
}

func (mr *mockCharacterRepository) GetCharacterById(id int64) (*entity.Character, error) {
	log.Printf("REPO MOCK: Get Character with id %d", id)
	arg := mr.Called(id)
	return arg.Get(0).(*entity.Character), arg.Error(1)
}

func (mr *mockCharacterRepository) InsertCharacter(character *entity.Character) error {
	log.Printf("REPO MOCK: Write Character %+v", character)
	arg := mr.Called(character)
	return arg.Error(0)
}

func TestCharacterService_FindCharacterById(t *testing.T) {
	var testCases = []struct {
		name     string
		id       int64
		response *entity.Character
		err      error
		// Repository
		repoLayer string
		repoRes   *entity.Character
		repoErr   error
	}{
		{
			"Should return 1 movie by id from MOCK repo",
			1,
			&entity.Character{ID: 1, Name: "Luz", Age: 14},
			nil,
			"mock",
			&testdata.Characters[0],
			nil,
		},
		{
			"Should return 1 movie by id from MOCKERY",
			1,
			&entity.Character{ID: 1, Name: "Luz", Age: 14},
			nil,
			"mockery",
			&testdata.Characters[0],
			nil,
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			// service initialize
			var svc CharacterService
			switch testCase.repoLayer {
			case "mock":
				repo := &mockCharacterRepository{}
				repo.On("GetCharacterById", testCase.id).Return(testCase.repoRes, testCase.repoErr)
				svc = NewCharacterService(repo)
			case "mockery":
				repo := mocks.NewCharacterRepository(t)
				repo.On("GetCharacterById", testCase.id).Return(testCase.repoRes, testCase.repoErr)
				svc = NewCharacterService(repo)
			case "integrated":
				repo := repository.NewCharacterRepository("fake-file")
				svc = NewCharacterService(repo)
			default:
				t.Fatalf("Should use valid repo: %v", testCase.repoLayer)
			}

			// Run test
			character, err := svc.FindCharacterById(testCase.id)
			t.Logf("Character found: %v", character)

			// Assert
			assert.Equal(t, testCase.response, character)
			assert.Equal(t, testCase.err, err)
		})
	}
}
