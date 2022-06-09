package service

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"toh-api/internal/entity"
	"toh-api/internal/service/mocks"
	"toh-api/test/testdata"
)

func TestCharacterService_FindCharacterById(t *testing.T) {
	var testCases = []struct {
		name     string
		id       int64
		response *entity.Character
		err      error
		// Repository
		repoRes *entity.Character
		repoErr error
	}{
		{
			"Should return 1 movie by id from MOCK repo",
			1,
			&entity.Character{ID: 1, Name: "Luz", Age: 14},
			nil,
			&testdata.Characters[0],
			nil,
		},
		{
			"Should return 1 movie by id from MOCKERY",
			1,
			&entity.Character{ID: 1, Name: "Luz", Age: 14},
			nil,
			&testdata.Characters[0],
			nil,
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			// service initialize
			var svc CharacterService

			repo := mocks.NewCharacterRepository(t)
			repo.On("GetCharacterById", testCase.id).Return(testCase.repoRes, testCase.repoErr)
			svc = NewCharacterService(repo)

			// Run test
			character, err := svc.FindCharacterById(testCase.id)
			t.Logf("Character found: %v", character)

			// Assert
			assert.Equal(t, testCase.response, character)
			assert.Equal(t, testCase.err, err)
		})
	}
}
