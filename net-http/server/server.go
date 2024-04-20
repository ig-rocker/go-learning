package server

import (
	"fmt"
	"net/http"
)

func Run() {

	http.Get("/")



	err := http.ListenAndServe(":8080",nil)
	// fmt.P
	if err!=nil{
		fmt.Println("server failed to start")
	}

}