package model

import "time"

type Employee struct {
	Name        string
	ID          string
	Designation string
	DOJ         time.Time
}


type EmployeesList struct{
	Data []Employee
}
