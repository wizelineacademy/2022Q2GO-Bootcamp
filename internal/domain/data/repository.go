package data

type Repository interface {
	Find(id int64) (*Data, error)
}
