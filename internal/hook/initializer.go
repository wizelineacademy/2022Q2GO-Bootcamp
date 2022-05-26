package hook

import (
	"fmt"
	"os"

	"github.com/krmirandas/2022Q2GO-Bootcamp/internal/model"
	"github.com/krmirandas/2022Q2GO-Bootcamp/internal/repository"
)

func getcwd() string {
	pwd, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	dataFilePath := fmt.Sprintf("%s/data.csv", pwd)

	return dataFilePath
}

var Pokemonss []model.Pokemon

func init() {
	//Pokemons
	fmt.Println(getcwd())
	data, err := repository.ReadCSVFromUrl(getcwd())

	if err != nil {
		fmt.Println(err)
	}

	// I dont understand yet why a global var works with = and not with :=
	Pokemonss = append(data)

	fmt.Println("Loading", len(Pokemonss), "Pokemons")
}
