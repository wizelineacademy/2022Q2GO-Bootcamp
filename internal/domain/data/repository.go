package data

// Data Repository
type Repository interface {
	Find(id int64) (*Data, error)
}

// Coffee repository
type RepoCoffee interface {
	GetCoffee() ([]Coffee, error)
}
