package repository

import (
	"fmt"
	"os"
	"testing"
)

func TestCharacterRepository_GetAllCharacters(t *testing.T) {
	path, err := os.Getwd()
	fmt.Println(path)
	repo := NewCharacterRepository("../../data/toh.csv")
	characters, err := repo.GetAllCharacters("even", 10, 2)
	if err != nil {
		return
	}
	if len(characters) != 10 {
		t.Errorf("repo.GetAllCharacters(\"even\", 10, 2); len() got %d; want 10", len(characters))
	}
}
