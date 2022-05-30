package repository

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Index int

func (i Index) GetId() int {
	return int(i)
}

type IndexList []Index

func (i IndexList) ListAll() ([]Index, error){
	return i, nil
}

func newIndexList(vals []Index) Repository[Index]{
	return IndexList(vals)
}

func TestGetById(t *testing.T) {
	list := newIndexList([]Index{0, 2, 1})

	idx, err := GetById(list, 2)
	assert.Nil(t, err)
	assert.Equal(t, Index(2), idx)

	_, err = GetById(list, 3)
	assert.NotNil(t, err)
}

