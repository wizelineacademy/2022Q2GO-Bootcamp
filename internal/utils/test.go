package utils

type Test struct {
	Name       string
	FileName   string
	Url        string
	HttpStatus int
	Response   string

	PokemonID int

	ReadingType     string
	Items           string
	ItemsPerWorkers string

	ExpectsError bool
}
