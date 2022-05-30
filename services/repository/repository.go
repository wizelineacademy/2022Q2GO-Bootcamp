package repository

import "errors"

type Indexable interface {
	GetId() int
}

type Repository[T Indexable] interface {
	ListAll() ([]T, error)
}

func GetById[T Indexable](r Repository[T], id int) (result T, err error) {
	elems, err := r.ListAll()

	if err != nil {
		return
	}

	for _, result = range elems {
		if result.GetId() == id {
			return
		}
	}
	return result, errors.New("ID not found")
}
