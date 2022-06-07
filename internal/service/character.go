package service

import (
	"toh-api/internal/entity"
)

// CharacterService the contract of the character service
type CharacterService interface {
	// CreateCharacter creates new character record
	CreateCharacter(character *entity.Character) error

	// FindCharacterById gets a character record by id
	FindCharacterById(id int64) (*entity.Character, error)
}

// CharacterRepository mocks dependency injection
type CharacterRepository interface {
	InsertCharacter(character *entity.Character) error
	GetCharacterById(id int64) (*entity.Character, error)
}

type characterService struct {
	repo CharacterRepository
}

func NewCharacterService(repo CharacterRepository) CharacterService {
	return &characterService{repo}
}

// - - - IMPLEMENTATION - - -

func (cs *characterService) CreateCharacter(character *entity.Character) error {
	err := cs.repo.InsertCharacter(character)
	if err == nil {
		return err
	}
	return nil
}

func (cs *characterService) FindCharacterById(id int64) (*entity.Character, error) {
	character, err := cs.repo.GetCharacterById(id)
	if err != nil {
		return nil, err
	}
	return character, nil
}
