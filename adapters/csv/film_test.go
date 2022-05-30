package csv

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFilmStructurer(t *testing.T) {
	valid := map[string]string{
		"id":           "1",
		"title":        "Wizeline Adventures",
		"director":     "Gopher",
		"release_date": "2022",
	}
	structurer := filmStructurer{}
	film, err := structurer.ToStruct(valid)
	assert.Nil(t, err)
	assert.NotNil(t, film)

	assert.Equal(t, film.Id, 1, "incorrect id parsing")
	assert.Equal(t, film.Title, "Wizeline Adventures", "incorrect title")

	invalid := map[string]string{
		"id":           "hola",
		"title":        "Wizeline Adventures",
		"director":     "Gopher",
		"release_date": "2022",
	}

	_, err = structurer.ToStruct(invalid)
	assert.NotNil(t, err)

	incomplete := map[string]string{
		"id":           "1",
		"director":     "Gopher",
		"release_date": "2022",
	}

	film, err = structurer.ToStruct(incomplete)
	assert.Nil(t, err)
	assert.Equal(t, film.Title, "", "title doesn't default to zero value")
}
