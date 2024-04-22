package handler

import (
	"net/http"
	model "test5/Model"
	"test5/Service"

	"github.com/gin-gonic/gin"
)


type IHandler interface { 
    AddEmployee(*gin.Context) 
    GetEmployee(*gin.Context) 
} 


type Handler struct{
    service.IService
}

func (i *Handler)AddEmployee(c *gin.Context){
    var reqBody model.Request

    if err:=c.BindJSON(&reqBody);err!=nil{
        c.JSON(http.StatusBadRequest, gin.H{"error: ":err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message":"Employee added Sucessfully"})
}

func (i *Handler)GetEmployee(c *gin.Context){
    data, err:= i.IService.GetEmployee("121")

    if err !=nil{
        c.JSON(http.StatusInternalServerError, gin.H{"error":err.Error()})
        return
    }

    c.JSON(http.StatusOK, data)

}


func NewHandler(s service.IService) IHandler { 
    return &Handler{
        IService :s,
    }
    //this should return implementation of IHandler 
} 

