package main

import (
	"net-http-basic/datastore"
	"net-http-basic/handler"
	"net-http-basic/service"
	"net/http"
)

func main() {
	data:=datastore.NewDatastore()

	s:=service.NewService(data)

	h:=handler.NewHandler(s)

	http.HandleFunc("/employee-list", h.HandleGetEmployeeList)

	http.HandleFunc("/employee",h.HandleGetEmployeeByID)

	http.HandleFunc("/create-employee",h.HandleCreateEmployee)

	http.HandleFunc("/update-employee",h.HandleUpdateEmployee)

	http.HandleFunc("/delete-employee",h.HandleDeleteEmployee)

	http.ListenAndServe(":8080", nil)

}
