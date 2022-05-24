package hooks

import (
	"fmt"

	"github.com/krmirandas/2022Q2GO-Bootcamp/api/model"
	"github.com/krmirandas/2022Q2GO-Bootcamp/api/repositories"
)

const dataFilePath = "/home/krmirandas/Documentos/Proyectos/Wizeline/Proyecto/2022Q2GO-Bootcamp/api/data.csv"

var Pokemonss []model.Pokemon

func init() {
	//Pokemons
	data, err := repositories.ReadCSVFromUrl(dataFilePath)

	if err != nil {
		fmt.Println(err)
	}

	// I dont understand yet why a global var works with = and not with :=
	Pokemonss = append(data)

	fmt.Println("Loading", len(Pokemonss), "Pokemons")
}
