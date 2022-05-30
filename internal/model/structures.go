package model

type ResponseError struct {
	Name string `json:"name,omitempty"`
}

type Pokemon struct {
	ID         string `json:"id"`
	Name       string `json:"name,omitempty"`
	Type1      string `json:"type1,omitempty"`
	Type2      string `json:"type2,omitempty"`
	Total      string `json:"total"`
	HP         string `json:"hp"`
	Attack     string `json:"attack"`
	Defense    string `json:"defense"`
	SpAtk      string `json:"spatk"`
	SpDef      string `json:"spdef"`
	Speed      string `json:"speed"`
	Generation string `json:"generation"`
	Legendary  string `json:"legendary"`
}

type ZipCodeInfo struct {
	PostCode            string                   `json:"post code,omitempty"`
	Country             string                   `json:"country,omitempty"`
	CountryAbbreviation string                   `json:"country abbreviation,omitempty"`
	Places              []map[string]interface{} `json:"places,omitempty"`
}
