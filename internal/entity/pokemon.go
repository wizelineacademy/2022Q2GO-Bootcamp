package entity

type Pokemon struct {
	Id   int    `json:"-"`
	Name string `json:"name"`
}
