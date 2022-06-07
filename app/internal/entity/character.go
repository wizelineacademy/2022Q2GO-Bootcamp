package models

type Origin struct {
	OName string `json:"name"`
	OUrl  string `json:"url"`
}

type Location struct {
	LName string `json:"name"`
	LUrl  string `json:"url"`
}

type Character struct {
	Id       int      `json:"id"`
	Name     string   `json:"name"`
	Status   string   `json:"status"`
	Species  string   `json:"species"`
	Type     string   `json:"type"`
	Gender   string   `json:"gender"`
	Image    string   `json:"image"`
	Episode  []string `json:"episode"`
	Url      string   `json:"url"`
	Created  string   `json:"created"`
	Origin   []Origin
	Location []Location
}

type CharacterDB struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	Status  string `json:"status"`
	Species string `json:"species"`
	Type    string `json:"type"`
	Gender  string `json:"gender"`
	Image   string `json:"image"`
	Url     string `json:"url"`
	Created string `json:"created"`
}

type Info struct {
	Count int         `json:"count"`
	Pages int         `json:"pages"`
	Next  string      `json:"next"`
	Prev  interface{} `json:"prev"`
}

type CharacterResponse struct {
	Results []Character
	Info    []Info
}
