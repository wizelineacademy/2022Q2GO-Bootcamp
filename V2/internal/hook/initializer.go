package hook

import (
	"fmt"
	"os"

	"github.com/krmirandas/2022Q2GO-Bootcamp/internal/model"
	"github.com/krmirandas/2022Q2GO-Bootcamp/internal/repository"
)

func Getcwd() string {
	pwd, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	dataFilePath := fmt.Sprintf("%s/data/data.csv", pwd)

	return dataFilePath
}

var Pokemonss []model.Pokemon

func init() {
	//Pokemons
	fmt.Println(Getcwd())
	data, err := repository.ReadCSVFromUrl(Getcwd())

	if err != nil {
		fmt.Println(err)
	}

	// I dont understand yet why a global var works with = and not with :=
	Pokemonss = append(data)

	fmt.Println("Loading", len(Pokemonss), "Pokemons")
}
