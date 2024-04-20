package service

import ("test5/Model"
    "test5/Datastore"
)

type Service struct{
    datastore Datastore.IDataStore
}

func (s *Service)AddEmployee(data model.Request)error{
    err:=s.datastore.Add(data)
    return err
}

func (s *Service)GetEmployee(id string)(model.Request,error){
    data, err:=s.datastore.Get(id)
    if err!=nil{
        return model.Request{},err
    }
    return data, nil
}



type IService interface { 
    AddEmployee(model.Request) error 
    GetEmployee(string) (model.Request, error) 
} 

func NewService(ds Datastore.IDataStore) IService { 
    //this should return implementation of IService 
    return &Service{
        datastore: ds,
    }
} 