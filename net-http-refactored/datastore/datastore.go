package datastore

import "customError"
type IBookstore interface{
	AddBook(models.Book)error
	GetBookCollections()models.BookCollection
	GetBookByID(string)models.Book
	UpdateBook(models.Book)(models.Book,error)
	DeleteBook(string)error
}


type Bookstore struct{
	models.BookCollection
}

func (bs *Bookstore)AddBook(b models.Book)(error){
	if _,ok:=bs.BookCollection[b.ID];ok{
		err:= customError.BookAlreadyExist{Message:"Already ID exist"}
		return err.Error()
	}
	bs.BookCollection[b.ID]=b
	return nil
}

func (bs *Bookstore)GetBookCollections()models.BookCollection{
	return bs.BookCollection
}

func(bs *Bookstore)GetBookByID(id string)models.Book{

	value, ok:=bs.BookCollection[]

}

func NewBookstore(books BookCollection)IBookstore{
	return &Bookstore{
		BookCollection: books,
	}
}