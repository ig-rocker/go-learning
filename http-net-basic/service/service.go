package service

import (
	"net-http-basic/datastore"
	"net-http-basic/model"
)

type Service struct {
	ds datastore.Datastore
}

func (s *Service) GetEmployesList() model.EmployeesList {
	return s.ds.GetEmployeeList()
}

func (s *Service) AddEmployee(employee model.Employee) error {
	err := s.ds.AddEmployee(employee)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) GetEmployeeListByID(id string)( model.Employee, error) {
	data,err := s.ds.GetEmployeeListByID(id)
	if err!=nil{
		return model.Employee{},err
	}
	return data,nil
}

func (s *Service) DeleteEmployee(id string) error {
	if err := s.ds.DeleteEmployee(id); err != nil {
		return err
	}
	return nil
}

func (s *Service) UpdateEmployee(employee model.Employee) error {
	if err := s.ds.UpdateEmployee(employee); err != nil {
		return err
	}
	return nil
}

func NewService(data datastore.Datastore) Service {
	return Service{
		ds: data,
	}
}
