package datastore

import (
	"errors"
	"net-http-basic/model"
)

type Datastore struct {
	data model.EmployeesList
}

// type Idatastore interface {
// 	GetEmployeeList() EmployeesList
// 	AddEmployee(model.Employee) error
// 	GetEmployeeListByID(string)model.Employee
// 	DeleteEmployee(string)error
// 	UpdateEmployee(string)model.Employee
// }

func (list *Datastore) GetEmployeeList() model.EmployeesList {
	return list.data
}

func (list *Datastore) AddEmployee(employee model.Employee) error {
	list.data.Data = append(list.data.Data, employee)
	return nil
}

func (list *Datastore) GetEmployeeListByID(id string) (model.Employee,error ){
	for _, value := range list.data.Data {
		if value.ID == id {
			return value, nil
		}
	}
	return model.Employee{},errors.New("ID not found")
}

func (list *Datastore) DeleteEmployee(id string) error {
	index := -1
	for x, value := range list.data.Data {
		if value.ID == id {
			index = x
			break
		}
	}
	if index > -1 {
		list.data.Data = append(list.data.Data[:index], list.data.Data[index+1:]...)
		return nil
	} else {
		return errors.New("ID not found")
	}
}

func (list *Datastore) UpdateEmployee(employee model.Employee) error {
	index := -1
	for x, value := range list.data.Data {
		if value.ID == employee.ID {
			index = x
			break
		}
	}
	if index > -1 {
		list.data.Data[index] = employee
		return nil
	} else {
		return errors.New("unable to find the element")
	}
}

func NewDatastore() Datastore {
	DefaultEmployeeList()

	return Datastore{
		// data:model.EmployeesList{Data:[]model.Employee{}},
		data: model.EmployeesList{Data: DataList.Data},
	}
}
