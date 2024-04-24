package customError

type BookAlreadyExist struct{
	Message string `json:"error"`
}

func (b BookAlreadyExist)Error()string{
	return b.Message
}