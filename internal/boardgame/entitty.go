package boardgame

type BoardGame struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	MinPlayers  string `json:"minPlayers"`
	MaxPlayers  string `json:"maxPlayers"`
	Duration    string `json:"duration"`
}
