package customError

import "fmt"

type BookAlreadyExist struct{
	Message string `json:"error"`
}

func (b BookAlreadyExist)Error()string{
	return b.Message
}

type BookNotExist struct{
	ID string
	Message string `json:"erorr"`
}

func (b BookNotExist)Error()string{
	 b.Message=fmt.Sprintf("ID %v not found",b.ID)
	 return b.Message
}