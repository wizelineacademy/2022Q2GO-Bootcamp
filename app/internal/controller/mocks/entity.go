package mocks

import (
	models "github.com/luischitala/2022Q2GO-Bootcamp/internal/entity"
)

// Mock Character
var Character = models.Character{
	Id:       0,
	Name:     "Rick Sanchez",
	Status:   "Alive",
	Species:  "Human",
	Type:     "",
	Gender:   "Male",
	Image:    "https://rickandmortyapi.com/api/character/avatar/1.jpeg",
	Episode:  nil,
	Url:      "https://rickandmortyapi.com/api/character/1",
	Created:  "2017-11-04T18:48:46.250Z",
	Origin:   nil,
	Location: nil,
}
