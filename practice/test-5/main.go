package main

import ("github.com/gin-gonic/gin"
	"test5/Datastore"
	"test5/Handler"
	"test5/Service"
)


func main() { 
    datastore := Datastore.NewDataStore()
    service := service.NewService(datastore)	
    handler := handler.NewHandler(service) 
    router := gin.Default() 
    router.POST("/employee", handler.AddEmployee) 
    router.GET("/employee/{id}", handler.GetEmployee) 
    router.Run("localhost:8080") 
}