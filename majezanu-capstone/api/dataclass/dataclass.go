package dataclass

type PokemonDataclass struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type ErrorDataclass struct {
	Code        int    `json:"code"`
	Description string `json:"description"`
}
