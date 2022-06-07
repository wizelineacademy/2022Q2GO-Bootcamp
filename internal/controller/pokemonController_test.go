package controller

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"github.com/McAdam17/2022Q2GO-Bootcamp/internal/repository"
	"github.com/McAdam17/2022Q2GO-Bootcamp/internal/service"
	"github.com/McAdam17/2022Q2GO-Bootcamp/internal/utils"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestHandleGetAllPokemons(t *testing.T) {
	allPokemonsOnCSV := `[{"id":1,"name":"bulbasaur"},{"id":2,"name":"ivysaur"},{"id":3,"name":"venusaur"}]`
	allPokemonsOnCSV = allPokemonsOnCSV + "\n"

	testCases := []utils.Test{
		{
			Name:         "test get all pokemons on csv file",
			ExpectsError: false,
			Url:          "/api/v1/all",
			FileName:     `/home/mcadam/Documents/platzi/wizeline/2022Q2GO-Bootcamp/pokemons_test.csv`,
			Response:     allPokemonsOnCSV,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.Name, func(t *testing.T) {
			testController := NewPokemonController(
				service.NewPokemonService(
					repository.NewPokemonRepository(testCase.FileName),
				),
			)

			e := echo.New()
			req := httptest.NewRequest(http.MethodGet, testCase.Url, strings.NewReader(testCase.Response))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)

			err := testController.HandleGetAllPokemons(c)

			if !testCase.ExpectsError {
				if assert.NoError(t, err) {
					assert.Equal(t, http.StatusOK, rec.Code)
					assert.Equal(t, testCase.Response, rec.Body.String())
				}
			}
		})
	}

}

func TestHandleGetPokemonById(t *testing.T) {
	pokemonWithId1 := `{"id":1,"name":"bulbasaur"}`
	pokemonWithId1 = pokemonWithId1 + "\n"

	pokemonNotFound := `{"message":"pokemon not found"}`
	pokemonNotFound = pokemonNotFound + "\n"

	testCases := []utils.Test{
		{
			Name:         "test get pokemons with id 1",
			ExpectsError: false,
			Url:          "/api/v1",
			PokemonID:    1,
			HttpStatus:   http.StatusOK,
			FileName:     `/home/mcadam/Documents/platzi/wizeline/2022Q2GO-Bootcamp/pokemons_test.csv`,
			Response:     pokemonWithId1,
		},
		{
			Name:         "test get pokemons with an invalid id format",
			ExpectsError: true,
			Url:          "/api/v1",
			PokemonID:    0,
			FileName:     `/home/mcadam/Documents/platzi/wizeline/2022Q2GO-Bootcamp/pokemons_test.csv`,
		},
		{
			Name:         "test get pokemons with an id does not exists",
			ExpectsError: false,
			Url:          "/api/v1",
			PokemonID:    4,
			HttpStatus:   http.StatusNotFound,
			FileName:     `/home/mcadam/Documents/platzi/wizeline/2022Q2GO-Bootcamp/pokemons_test.csv`,
			Response:     pokemonNotFound,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.Name, func(t *testing.T) {
			testController := NewPokemonController(
				service.NewPokemonService(
					repository.NewPokemonRepository(testCase.FileName),
				),
			)

			e := echo.New()
			req := httptest.NewRequest(http.MethodGet, testCase.Url, strings.NewReader(testCase.Response))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)
			c.SetParamNames("id")
			if testCase.PokemonID != 0 {
				c.SetParamValues(fmt.Sprint(testCase.PokemonID))
			}

			err := testController.HandleGetPokemonById(c)

			if !testCase.ExpectsError {
				if assert.NoError(t, err) {
					assert.Equal(t, testCase.HttpStatus, rec.Code)
					assert.Equal(t, testCase.Response, rec.Body.String())
				}
			}
		})
	}

}

func TestHandleGetPokemonItemsFromCSV(t *testing.T) {
	invalidType := `{"message":"invalid type"}`
	invalidType = invalidType + "\n"
	invalidItems := `{"message":"invalid items"}`
	invalidItems = invalidItems + "\n"
	invalidWorkers := `{"message":"invalid items_per_workers"}`
	invalidWorkers = invalidWorkers + "\n"
	oddPokemons := `[{"id":1,"name":"bulbasaur"},{"id":3,"name":"venusaur"}]`
	oddPokemons = oddPokemons + "\n"
	evenPokemons := `[{"id":2,"name":"ivysaur"},{"id":4,"name":"pikachu"}]`
	evenPokemons = evenPokemons + "\n"

	testCases := []utils.Test{
		{
			Name:         "test invalid type",
			ExpectsError: false,
			Url:          "/api/v1/add/generation",
			HttpStatus:   http.StatusBadRequest,
			FileName:     `/home/mcadam/Documents/platzi/wizeline/2022Q2GO-Bootcamp/pokemons_test.csv`,
			Response:     invalidType,
		},
		{
			Name:         "test invalid items",
			ExpectsError: false,
			Url:          "/api/v1/add/generation",
			HttpStatus:   http.StatusBadRequest,
			ReadingType:  "odd",
			FileName:     `/home/mcadam/Documents/platzi/wizeline/2022Q2GO-Bootcamp/pokemons_test.csv`,
			Response:     invalidItems,
		},
		{
			Name:         "test invalid workers",
			ExpectsError: false,
			Url:          "/api/v1/add/generation",
			HttpStatus:   http.StatusBadRequest,
			ReadingType:  "odd",
			Items:        "2",
			FileName:     `/home/mcadam/Documents/platzi/wizeline/2022Q2GO-Bootcamp/pokemons_test.csv`,
			Response:     invalidWorkers,
		},
		{
			Name:            "test get valid 2 first odd pokemons",
			ExpectsError:    false,
			Url:             "/api/v1/add/generation",
			HttpStatus:      http.StatusOK,
			ReadingType:     "odd",
			Items:           "2",
			ItemsPerWorkers: "1",
			FileName:        `/home/mcadam/Documents/platzi/wizeline/2022Q2GO-Bootcamp/pokemons_test_2.csv`,
			Response:        oddPokemons,
		},
		{
			Name:            "test get valid 2 first even pokemons",
			ExpectsError:    false,
			Url:             "/api/v1/add/generation",
			HttpStatus:      http.StatusOK,
			ReadingType:     "even",
			Items:           "2",
			ItemsPerWorkers: "1",
			FileName:        `/home/mcadam/Documents/platzi/wizeline/2022Q2GO-Bootcamp/pokemons_test_2.csv`,
			Response:        evenPokemons,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.Name, func(t *testing.T) {
			testController := NewPokemonController(
				service.NewPokemonService(
					repository.NewPokemonRepository(testCase.FileName),
				),
			)

			e := echo.New()
			q := make(url.Values)
			q.Set("type", testCase.ReadingType)
			q.Set("items", testCase.Items)
			q.Set("items_per_workers", testCase.ItemsPerWorkers)
			req := httptest.NewRequest(http.MethodGet, testCase.Url+"?"+q.Encode(), strings.NewReader(testCase.Response))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)

			err := testController.HandleGetPokemonItemsFromCSV(c)

			if !testCase.ExpectsError {
				if assert.NoError(t, err) {
					assert.Equal(t, testCase.HttpStatus, rec.Code)
					assert.Equal(t, testCase.Response, rec.Body.String())
				}
			}
		})
	}

}
