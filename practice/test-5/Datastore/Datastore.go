package Datastore

import (
	"errors"
	model "test5/Model"
)

type DataStore struct {
	Data map[string]model.Request
}
type IDataStore interface {
	Add(model.Request) error
	Get(string) (model.Request, error)
}

func (ds *DataStore) Add(data model.Request) error {
	ds.Data[data.Id] = data
	return nil
}

func (ds *DataStore) Get(key string) (model.Request, error) {
	req, ok := ds.Data["key"]
	if !ok {
		return model.Request{}, errors.New("Unable to find the error")
	}
	return req, nil
}

func NewDataStore() IDataStore {
	//this should return implementation of IDataStore
	return &DataStore{
		Data: make(map[string]model.Request),
	}
}
