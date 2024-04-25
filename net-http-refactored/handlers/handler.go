package handlers

import("services")

type IHandler interface{
	AddBook(models.Book)error
	GetBookCollection()models.BookCollection
	GetBookByID(string)(models.Book, error)
}

type Handler struct{
	services.IService
}


func (h *Handler)AddBook(w http.WriteHeader, r *http.Request){
	
}

