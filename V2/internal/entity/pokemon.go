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
	ID                     string `json:"id" validate:"required"`
	Name                   string `json:"name" validate:"required"`
	BaseExperience         string `json:"base_experience" validate:"required"`
	Height                 string `json:"is_default" validate:"required"`
	IsDefault              string `json:"total" validate:"required"`
	Order                  string `json:"order" validate:"required"`
	Weight                 string `json:"weight" validate:"required"`
	Abilities              string `json:"abilities" validate:"required"`
	Forms                  string `json:"forms" validate:"required"`
	GameIndices            string `json:"game_indices" validate:"required"`
	HeldItems              string `json:"held_items" validate:"required"`
	LocationAreaEncounters string `json:"location_area_encounters" validate:"required"`
	PastTypes              string `json:"past_types" validate:"required"`
	Sprites                string `json:"sprites" validate:"required"`
	Species                string `json:"species" validate:"required"`
	Stats                  string `json:"stats" validate:"required"`
	Types                  string `json:"types" validate:"required"`
}
