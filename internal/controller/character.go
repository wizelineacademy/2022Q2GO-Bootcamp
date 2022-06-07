package controller

//
//import (
//	"github.com/gofiber/fiber/v2"
//	"toh-api/internal/entity"
//)
//
//// CharacterController the contract of the movie controller
//type CharacterController interface {
//	// CreateCharacter creates a new character
//	CreateCharacter(ctx *fiber.Ctx)
//
//	// FindCharacter gets a character by id
//	FindCharacter(ctx *fiber.Ctx)
//
//	// UpdateCharacter updates the record character
//	UpdateCharacter()
//
//	// DeleteCharacter deletes one character
//	DeleteCharacter()
//}
//
//// CharacterService service dependency injection
//type CharacterService interface {
//	// CreateCharacter creates new character record
//	CreateCharacter(character *entity.Character) error
//
//	// FindCharacterById gets a character record by id
//	FindCharacterById(id int64) (*entity.Character, error)
//}
//
//type characterController struct {
//	svc CharacterService
//}
//
//func NewCharacterController(svc CharacterService) CharacterController {
//	return &characterController{svc: svc}
//}
//
//// IMPLEMENTATION
//
//func (cc *characterController) FindCharacter(ctx *fiber.Ctx) {
//
//}
//
//func (cc *characterController) CreateCharacter(ctx *fiber.Ctx) {
//
//}
