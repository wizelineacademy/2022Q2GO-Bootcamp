package routes

import (
	"github.com/gofiber/fiber/v2"
	"net/http"
	"strconv"
	"toh-api/internal/repository"
	"toh-api/internal/service"
	"toh-api/pkg/parser"
)

type characterRoutes struct{}

func CharacterRoutes() characterRoutes {
	return characterRoutes{}
}

// swagger:route GET /character/:id characterById
// Returns a character based on the given id
// responses:
// 	201: character

// getCharacter returns the character from the data store
func (u characterRoutes) getCharacter(ctx *fiber.Ctx) error {
	id := ctx.Params("id", "")

	idInt, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return fiber.NewError(http.StatusBadRequest, "ID should be an INT!")
	}

	repo := repository.NewCharacterRepository("./data/toh.csv")
	charService := service.NewCharacterService(repo)

	character, err := charService.FindCharacterById(idInt)
	if err != nil {
		return fiber.NewError(http.StatusNotFound, "Record not found!")
	}

	return ctx.Status(http.StatusCreated).JSON(character)
}

func (u characterRoutes) RegisterRoutes(api *parser.ApiService) {
	api.GetPublic("/character/:id", u.getCharacter)
}
