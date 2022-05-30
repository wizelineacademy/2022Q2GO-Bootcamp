package entities

type Film struct {
	Id int `json:"id"`
	Title string `json:"title"`
	Director string `json:"director"`
	ReleaseDate string `json:"release_date"`
}

func (f Film) GetId() int {
	return f.Id
}
