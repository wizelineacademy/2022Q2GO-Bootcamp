package controllers

import (
	"encoding/json"
	"io"
	"net/http/httptest"
	"strings"
	"testing"
	"wizeline/ghibli/entities"
	"wizeline/ghibli/services/repository"

	"github.com/stretchr/testify/assert"
)

func TestSendEncoded(t *testing.T) {
	str := "hello gophers!"
	data, _ := json.Marshal(str)
	writer := httptest.NewRecorder()
	sendEncoded(writer, data)

	assert.Equal(t, "application/json", writer.Header().Get("Content-Type"))

	codedData, err := io.ReadAll(writer.Result().Body)
	assert.Nil(t, err)

	var resp string
	err = json.Unmarshal(codedData, &resp)
	assert.Nil(t, err)
	assert.Equal(t, str, resp)
}

type FilmList []entities.Film

func (f FilmList) ListAll() ([]entities.Film, error) {
	return f, nil
}

func newFilmList(films []entities.Film) repository.Repository[entities.Film] {
	return FilmList(films)
}

func TestGetEncodedFilms(t *testing.T) {
	film := entities.Film{Id: 1, Title: "My Awesome Movie"}
	controller := NewFilmController(newFilmList([]entities.Film{
		film,
	}))

	data, err := controller.getEncodedFilm("1")
	assert.Nil(t, err)

	var respFilm entities.Film
	err = json.Unmarshal(data, &respFilm)
	assert.Nil(t, err)
	assert.Equal(t, film, respFilm)

	data, err = controller.getEncodedFilms()
	assert.Nil(t, err)

	var respFilms []entities.Film
	err = json.Unmarshal(data, &respFilms)
	assert.Nil(t, err)
	assert.Equal(t, []entities.Film{film}, respFilms)
}

func TestTryGetEncoded(t *testing.T) {
	film := entities.Film{Id: 1, Title: "My Awesome Movie"}
	controller := NewFilmController(newFilmList([]entities.Film{
		film,
	}))

	req := httptest.NewRequest(
		"GET",
		"http://localhost:8009/films?id=1",
		strings.NewReader(""),
	)

	data, err := controller.tryGetEncoded(req)
	assert.Nil(t, err)

	var respFilm entities.Film
	err = json.Unmarshal(data, &respFilm)
	assert.Nil(t, err)
	assert.Equal(t, film, respFilm)

	req = httptest.NewRequest(
		"GET",
		"http://localhost:8009/films",
		strings.NewReader(""),
	)

	data, err = controller.tryGetEncoded(req)
	assert.Nil(t, err)

	var respFilms []entities.Film
	err = json.Unmarshal(data, &respFilms)
	assert.Nil(t, err)
	assert.Equal(t, []entities.Film{film}, respFilms)
}
