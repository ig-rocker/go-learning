package server

import (
	"fmt"
	"net/http"
)

func Run() {


	http.HandleFunc("/",(w http.Writer, r *http.Request){
		w.WriteHeader(http.StatusOK)
		w.Write([]byte{"Welcome to Home Page"})
		return
	})



	err := http.ListenAndServe(":8080",nil)
	// fmt.P
	if err!=nil{
		fmt.Println("server failed to start")
	}

}