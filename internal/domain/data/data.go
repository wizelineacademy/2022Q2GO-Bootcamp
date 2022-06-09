package data

// Data model
type Data struct {
	Id   int64
	Item string
}

// Coffee model
type Coffee struct {
	Id          int64  `json:"id"`
	Name        string `json:"blend_name"`
	Description string `json:"notes"`
	Origin      string `json:"origin"`
	Variety     string `json:"variety"`
}
