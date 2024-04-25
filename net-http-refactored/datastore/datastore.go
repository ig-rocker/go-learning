package datastore

import "customError"
type IBookstore interface{
	Add(models.Book)error
	Getcollection()models.BookCollection
	GetID(string)models.Book
	// UpdateBook(models.Book)(models.Book,error)
	// DeleteBook(string)error
}


type Bookstore struct{
	models.BookCollection
}

func (bs *Bookstore)Add(b models.Book)(error){
	if _,ok:=bs.BookCollection[b.ID];ok{
		err:= customError.BookAlreadyExist{Message:"Already ID exist"}
		return err
	}
	bs.BookCollection[b.ID]=b
	return nil
}

func (bs *Bookstore)Getcollections()models.BookCollection{
	return bs.BookCollection
}

func(bs *Bookstore)GetID(id string)(models.Book, error){

	value, ok:=bs.BookCollection.collection[id]
	if !ok{
		err:=customError.BookNotExist{ID:id}
		return nil,err
	}
	return value, nil

}

func NewBookstore(books BookCollection)IBookstore{
	return &Bookstore{
		BookCollection: books,
	}
}