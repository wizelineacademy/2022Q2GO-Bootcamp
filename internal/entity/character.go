package entity

// A Character from TOH.
// swagger:response character
type Character struct {
	// Example: 1
	ID int64 `json:"id"`
	// Required: true
	// Example: Luz
	Name string `json:"name"`
	// Required: true
	// Example: 14
	Age int64 `json:"age"`
}
