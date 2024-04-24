package datastore

import (
	"net-http-basic/model"
	"time"
)

var DataList = model.EmployeesList{
	Data: []model.Employee{},
}

func DefaultEmployeeList() {
	DataList.Data = append(DataList.Data, model.Employee{"John", "100", "Manager", time.Now()})
	DataList.Data = append(DataList.Data, model.Employee{"Mayer", "101", "Lead", time.Now()})
	DataList.Data = append(DataList.Data, model.Employee{"Mia", "102", "PA", time.Now()})
	DataList.Data = append(DataList.Data, model.Employee{"Dani", "103", "HR", time.Now()})
}
