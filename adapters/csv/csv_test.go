package csv

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Stringer struct {}

func (s Stringer) ToStruct(values map[string]string) (string, error) {
	return values["string"], nil
}

func newStringer() Structurer[string] {
	return Stringer{}
}

func TestCsvRepo(t *testing.T) {
	reader := strings.NewReader("string\ngolang\nis\nawesome")
	repo := NewCsvRepository(reader, newStringer())
	assert.NotNil(t, repo)

	strs, err := repo.ListAll()
	assert.Nil(t, err)
	assert.Equal(t, []string{"golang", "is", "awesome"}, strs)
}
