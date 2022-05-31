package query

import (
	"github.com/jesusrevilla/capstone/internal/domain/data"
)

// Get Data Request Model of the handler
type DataRequest struct {
	DataId int64
}

// Get Data Result is the return model of Data Query Handler
type DataResult struct {
	Id   int64
	Item string
}

// DataRequestHandler provides an interface to handle a DataRequest and return a *DataResult
type DataRequestHandler interface {
	Handle(query DataRequest) (*DataResult, error)
}

type dataRequestHandler struct {
	repo data.Repository
}

// Handler Constructor
func NewDataRequestHandler(repo data.Repository) DataRequestHandler {
	return dataRequestHandler{repo: repo}
}

// Handles DataRequest query
func (h dataRequestHandler) Handle(query DataRequest) (*DataResult, error) {
	data, err := h.repo.Find(query.DataId)
	var result *DataResult
	if data != nil && err == nil {
		result = &DataResult{Id: data.Id, Item: data.Item}
	}
	return result, err
}
