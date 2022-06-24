package entity

type Pokemon struct {
	ID         string `json:"id" validate:"required" validate:"required"`
	Name       string `json:"name" validate:"required"`
	Type1      string `json:"type1" validate:"required"`
	Type2      string `json:"type2" validate:"required"`
	Total      string `json:"total" validate:"required"`
	HP         string `json:"hp" validate:"required"`
	Attack     string `json:"attack" validate:"required"`
	Defense    string `json:"defense" validate:"required"`
	SpAtk      string `json:"spatk" validate:"required"`
	SpDef      string `json:"spdef" validate:"required"`
	Speed      string `json:"speed" validate:"required"`
	Generation string `json:"generation" validate:"required"`
	Legendary  string `json:"legendary" validate:"required"`
}

type PokemonInfo struct {
	ID                     int    `json:"id" validate:"required"`
	Name                   string `json:"name" validate:"required"`
	BaseExperience         int    `json:"base_experience" validate:"required"`
	Height                 int    `json:"height" validate:"required"`
	IsDefault              bool   `json:"is_default" validate:"required"`
	Order                  int    `json:"order" validate:"required"`
	Weight                 int    `json:"weight" validate:"required"`
	LocationAreaEncounters string `json:"location_area_encounters" validate:"required"`
}
