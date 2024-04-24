package handler

import (
	"encoding/json"
	"net-http-basic/model"
	"net-http-basic/service"
	"net/http"
)

type Handler struct {
	service.Service
}

func (h *Handler) HandleGetEmployeeList(w http.ResponseWriter, r *http.Request) {
	if r.Method=="GET"{
	data := h.Service.GetEmployesList()
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(data.Data)
	}else{
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Bad Request"))
	}
}

func (h *Handler) HandleGetEmployeeByID(w http.ResponseWriter, r *http.Request) {
	if r.Method=="GET"{
	id := r.URL.Query().Get("id")
	data, err := h.Service.GetEmployeeListByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(data)
	}else{
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Bad Request"))	
	}
}

func (h *Handler) HandleCreateEmployee(w http.ResponseWriter, r *http.Request) {
	if r.Method=="POST"{
		var emp model.Employee

		if err:= json.NewDecoder(r.Body).Decode(&emp);err!=nil{
			http.Error(w,err.Error(),http.StatusBadRequest)
			return
		}

		err:=h.Service.AddEmployee(emp)
		if err!=nil{
			http.Error(w, err.Error(),http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(emp)
	}else{
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Bad Request"))
	}

}

func (h *Handler) HandleUpdateEmployee(w http.ResponseWriter, r *http.Request) {
	if r.Method=="PUT"{
		w.Header().Set("Content-Type","application/json")
			var employee model.Employee
			err:=json.NewDecoder(r.Body).Decode(&employee)
			if err!=nil{
				http.Error(w, err.Error(),http.StatusBadRequest)
				return
			}
	
			if err:=h.Service.UpdateEmployee(employee); err!=nil{
				http.Error(w, err.Error(),http.StatusInternalServerError)
				return
			}
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(&employee)
	}else{
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Bad Request"))
	}
}

func (h *Handler) HandleDeleteEmployee(w http.ResponseWriter, r *http.Request) {
	if r.Method=="DELETE"{
		id:=r.URL.Query().Get("id")
		if err:=h.Service.DeleteEmployee(id);err!=nil{
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Succesfully Deleted"))
	}else{
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Bad Request"))
	}

}

func NewHandler(s service.Service) Handler {
	return Handler{
		Service: s,
	}
}
