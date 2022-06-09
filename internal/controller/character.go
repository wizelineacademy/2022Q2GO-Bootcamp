package controller

import (
	"github.com/gofiber/fiber/v2"
	"net/http"
	"strconv"
	"toh-api/internal/entity"
)

// CharacterController the contract of the movie controller
type CharacterController interface {
	// CreateCharacter creates a new character
	CreateCharacter(ctx *fiber.Ctx) error

	// FindCharacter gets a character by id
	FindCharacter(ctx *fiber.Ctx) error

	// UpdateCharacter updates the record character
	//UpdateCharacter()

	// DeleteCharacter deletes one character
	//DeleteCharacter()
}

// CharacterService service dependency injection
type CharacterService interface {
	// CreateCharacter creates new character record
	CreateCharacter(character *entity.Character) error

	// FindCharacterById gets a character record by id
	FindCharacterById(id int64) (*entity.Character, error)
}

type characterController struct {
	svc CharacterService
}

// NewCharacterController constructor for characterController
func NewCharacterController(svc CharacterService) CharacterController {
	return &characterController{svc}
}

// IMPLEMENTATION

// FindCharacter returns an existant character
func (cc *characterController) FindCharacter(ctx *fiber.Ctx) error {
	id := ctx.Params("id", "")

	idInt, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return fiber.NewError(http.StatusBadRequest, "ID should be an INT!")
	}

	character, err := cc.svc.FindCharacterById(idInt)
	if err != nil {
		return fiber.NewError(http.StatusNotFound, "Record not found!")
	}

	return ctx.Status(http.StatusCreated).JSON(character)
}

// CreateCharacter creates a new character
func (cc *characterController) CreateCharacter(ctx *fiber.Ctx) error {
	// TODO: implement me
	panic("implement me")

	return nil
}
