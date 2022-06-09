package routes

import (
	"github.com/gofiber/fiber/v2"
	"toh-api/cmd/server/api"
	"toh-api/internal/controller"
	"toh-api/internal/repository"
	"toh-api/internal/service"
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
	repo := repository.NewCharacterRepository("./data/toh.csv")
	charService := service.NewCharacterService(repo)
	myController := controller.NewCharacterController(charService)

	return myController.FindCharacter(ctx)
}

// RegisterRoutes registers the routes related to characters
func (u characterRoutes) RegisterRoutes(api *api.ApiService) {
	api.GetPublic("/character/:id", u.getCharacter)
}
