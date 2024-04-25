package services

import (
	"models"
	"datastore"
)

type IService interface{
	AddBook(models.Book)error
	GetBookCollection()models.BookCollection
	GetBookByID(string)(models.Book, error)
}

type Service struct{
	datastore.IBookstore
}

func(s *Service)AddBook(models.Book)error{
	if err:=s.IBookstore.Add(models.Book);err!=nil{
		return err.Error()
	}
	return nil
}

func(s *Service)GetBookCollection()models.BookCollection{
	return s.IBookstore.Getcollections()
}

func (s *Service)GetBookByID(id string)(models.Book, error){
	data, err:=s.IBookstore.GetID(id)
	if err!=nil{
		return nil, err.Error()
	}
	return data, nil
}


func NewService(ds datastore.IBookstore)IService{
	return &Service{
		IBookstore:ds
	}
}



