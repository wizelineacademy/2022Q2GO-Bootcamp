package routes

import (
	"github.com/gofiber/fiber/v2"
	"net/http"
	"strconv"
	"toh-api/cmd/server/api"
	"toh-api/internal/repository"
	"toh-api/internal/service"
)

type characterRoutes struct{}

func CharacterRoutes() characterRoutes {
	return characterRoutes{}
}

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

func (u characterRoutes) RegisterRoutes(api *api.ApiService) {
	api.GetPublic("/character/:id", u.getCharacter)
}
